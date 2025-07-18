package user

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"time"

	reviewenum "golang-auth-app/app/common/enums/review"
	statusenum "golang-auth-app/app/common/enums/status"
	config "golang-auth-app/config"

	"golang-auth-app/app/adapters/sql/gorm/model"

	"golang-auth-app/app/common/errorcode"
	userDto "golang-auth-app/app/interfaces/management/user/dto"
	"golang-auth-app/app/interfaces/smtp"

	fileutil "golang-auth-app/app/utils/file"
	shautil "golang-auth-app/app/utils/sha"
	sliceutil "golang-auth-app/app/utils/slice"

	"github.com/rotisserie/eris"
)

func (i *impl) Review(
	ctx context.Context,
	payload *userDto.ServiceReviewPayload,
) error {
	var err error
	now := time.Now()

	toBeUpdatedUser := &model.User{
		ID:        payload.ExistingUser.ID,
		Status:    payload.ExistingUser.Status,
		UpdatedBy: payload.Modifier,
		UpdatedAt: now,
	}

	currentStatus := statusenum.Status(payload.ExistingUser.Status)
	var newStatus statusenum.Status
	var isNewAccount bool
	var activityLogMessage string

	switch payload.Action {
	case reviewenum.Approve:
		switch currentStatus {
		case statusenum.Submitted:
			newStatus = statusenum.Active
			toBeUpdatedUser.StartDate = &now
			isNewAccount = true
			activityLogMessage = fmt.Sprintf("User was approved by %s", payload.Modifier)
		case statusenum.ActiveEditSubmitted:
			newStatus = statusenum.Active
			activityLogMessage = fmt.Sprintf("User changes was approved by %s", payload.Modifier)
		case statusenum.ActiveInactivationSubmitted:
			newStatus = statusenum.Inactive
			toBeUpdatedUser.EndDate = &now
			activityLogMessage = fmt.Sprintf("User inactivation was approved by %s", payload.Modifier)
		}
	case reviewenum.Reject:
		switch currentStatus {
		case statusenum.Submitted:
			newStatus = statusenum.Rejected
			activityLogMessage = fmt.Sprintf("User was rejected by %s", payload.Modifier)
		case statusenum.ActiveEditSubmitted, statusenum.ActiveInactivationSubmitted:
			newStatus = statusenum.Active
			activityLogMessage = fmt.Sprintf("User changes was rejected by %s", payload.Modifier)
		}
	}
	toBeUpdatedUser.Status = newStatus.ToString()

	/*
		Flows
		1. action == approve, currentStatus == Submitted => newStatus = Active => copy userRoleMappings => notify new user account via email
		2. action == approve, currentStatus == ActiveEditSubmitted => newStatus = Active => copy userRoleMappings => delete tempUser
		3. action == approve, currentstatus == ActiveInactivationSubmitted => newStatus = Inactive => fill inactive date user
		4. action == reject, currentStatus == Submitted => newStatus = Rejected => change status from submitted into rejected
		5. action == reject, currentStatus == ActiveEditSubmitted => newStatus => Active => copy tempUser => delete tempUser
		6. action == reject, currentStatus = ActiveInactivationSubmitted => newStatus = Active
	*/

	switch reviewenum.Action(payload.Action) {
	case reviewenum.Approve:
		existingUserRoles, err := i.userSqlAdapter.GetUserRoleMappings(ctx, nil, []int32{payload.ExistingUser.ID})
		if err != nil {
			return err
		}

		switch newStatus {
		case statusenum.Active:
			var payloadAssignedRoles []*model.UserRoleMapping
			if err = json.Unmarshal([]byte(*payload.ExistingUser.AssignedRoles), &payloadAssignedRoles); err != nil {
				return eris.Wrap(err, err.Error())
			}

			err = i.deleteUserRoleMappings(ctx, existingUserRoles, payloadAssignedRoles, payload.Modifier)
			if err != nil {
				return err
			}

			err = i.insertNewUserRoleMappings(ctx, existingUserRoles, payloadAssignedRoles)
			if err != nil {
				return err
			}
		case statusenum.Inactive:
			err = i.userSqlAdapter.DeleteUserRoleMappings(ctx, sliceutil.Map(existingUserRoles, func(dt *model.UserRoleMapping) int32 {
				return dt.ID
			}), payload.Modifier)
			if err != nil {
				return err
			}
		}
	case reviewenum.Reject:
		switch currentStatus {
		case statusenum.ActiveEditSubmitted:
			existingTempUser, err := i.userSqlAdapter.GetTempUserByUserId(ctx, payload.ExistingUser.ID)
			if err != nil {
				return err
			} else if existingTempUser == nil {
				return errorcode.WithCustomMessage(errorcode.ErrCodeNotFound, "temp user is not found")
			}

			toBeUpdatedUser.Username = existingTempUser.Username
			toBeUpdatedUser.FullName = existingTempUser.FullName
			toBeUpdatedUser.Description = existingTempUser.Description
			toBeUpdatedUser.Email = existingTempUser.Email
			toBeUpdatedUser.PhoneNumber = existingTempUser.PhoneNumber
			toBeUpdatedUser.AssignedRoles = existingTempUser.AssignedRoles
		}
	}

	defaultPassword := rand.Text()
	if isNewAccount {
		toBeUpdatedUser.Password = shautil.EncryptString(defaultPassword)
	}

	err = i.userSqlAdapter.UpdateUser(ctx, toBeUpdatedUser)
	if err != nil {
		return err
	}

	if isNewAccount {
		err = i.sendEmailNewUser(ctx, payload.ExistingUser.Username, payload.ExistingUser.Email, defaultPassword)
		if err != nil {
			return err
		}
	}

	if activityLogMessage != "" && payload.ExistingUser.ActivityLogID != nil {
		err = i.activityLogSqlAdapter.Insert(ctx, *payload.ExistingUser.ActivityLogID, activityLogMessage, "SUCCESS")
		if err != nil {
			return err
		}
	}

	return nil
}

func (i *impl) deleteUserRoleMappings(
	ctx context.Context,
	existingData []*model.UserRoleMapping,
	payload []*model.UserRoleMapping,
	modifier string,
) error {
	payloadMapped := sliceutil.AssociateBy(payload, func(dt *model.UserRoleMapping) string {
		return fmt.Sprintf("%d:%d", dt.UserID, dt.RoleID)
	})

	toBeRemovedDataIds := []int32{}
	for _, dt := range existingData {
		if _, isExist := payloadMapped[fmt.Sprintf("%d:%d", dt.UserID, dt.RoleID)]; !isExist {
			toBeRemovedDataIds = append(toBeRemovedDataIds, dt.ID)
		}
	}

	if len(toBeRemovedDataIds) > 0 {
		if err := i.userSqlAdapter.DeleteUserRoleMappings(ctx, toBeRemovedDataIds, modifier); err != nil {
			return err
		}
	}

	return nil
}

func (i *impl) insertNewUserRoleMappings(
	ctx context.Context,
	existingData []*model.UserRoleMapping,
	payload []*model.UserRoleMapping,
) error {
	existingDataMapped := sliceutil.AssociateBy(existingData, func(dt *model.UserRoleMapping) string {
		return fmt.Sprintf("%d:%d", dt.UserID, dt.RoleID)
	})

	toBeInsertedData := []*model.UserRoleMapping{}
	for _, dt := range payload {
		if _, isExist := existingDataMapped[fmt.Sprintf("%d:%d", dt.UserID, dt.RoleID)]; !isExist {
			toBeInsertedData = append(toBeInsertedData, dt)
		}
	}

	if len(toBeInsertedData) > 0 {
		if err := i.userSqlAdapter.InsertUserRoleMappings(ctx, toBeInsertedData); err != nil {
			return err
		}
	}

	return nil
}

func (i *impl) sendEmailNewUser(ctx context.Context, username, email, defaultPassword string) error {
	loginUrl := config.Module.Auth.WebPageUrl.LoginUrl

	type htmlTemplate struct {
		Username        string
		Email           string
		DefaultPassword string
		LoginUrl        string
	}

	htmlBody, err := fileutil.ParseTemplateFile(ctx, "./etc/template/authentication/new-account-email.html", htmlTemplate{
		Username:        username,
		Email:           email,
		DefaultPassword: defaultPassword,
		LoginUrl:        loginUrl,
	})
	if err != nil {
		return err
	}

	cfgSMTP := config.ExternalService.Smtp
	if err = i.smtpAdapter.SendEmail(ctx, &smtp.SendEmailPayload{
		Host:       cfgSMTP.Host,
		Port:       cfgSMTP.Port,
		User:       cfgSMTP.User,
		Password:   cfgSMTP.Password,
		Recipients: []string{email},
		Subject:    "Welcome to Golauthapp - Your Account Has Been Created",
		HTMLBody:   htmlBody,
	}); err != nil {
		return err
	}

	return nil
}
