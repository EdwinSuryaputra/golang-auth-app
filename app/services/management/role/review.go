package role

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	reviewenum "golang-auth-app/app/common/enums/review"
	statusenum "golang-auth-app/app/common/enums/status"
	"golang-auth-app/app/datasources/sql/gorm/model"
	roleInterface "golang-auth-app/app/interfaces/management/role"

	sliceutil "golang-auth-app/app/utils/slice"

	"github.com/rotisserie/eris"
)

func (i *impl) Review(ctx context.Context, payload *roleInterface.ServiceReviewRolePayload) error {
	var err error
	now := time.Now()

	toBeUpdatedRole := &model.Role{
		ID:        payload.ExistingRole.ID,
		UpdatedBy: payload.Modifier,
		UpdatedAt: time.Now(),
	}

	currentStatus := statusenum.Status(payload.ExistingRole.Status)
	var newStatus statusenum.Status
	var activityLogMessage string

	switch payload.Action {
	case reviewenum.Approve:
		switch currentStatus {
		case statusenum.Submitted:
			newStatus = statusenum.Active
			activityLogMessage = fmt.Sprintf("Role was approved by %s", payload.Modifier)
		case statusenum.ActiveEditSubmitted:
			newStatus = statusenum.Active
			activityLogMessage = fmt.Sprintf("Role changes was approved by %s", payload.Modifier)
		case statusenum.ActiveInactivationSubmitted:
			newStatus = statusenum.Inactive
			toBeUpdatedRole.InactiveDate = &now
			activityLogMessage = fmt.Sprintf("Role inactivation was approved by %s", payload.Modifier)
		}
	case reviewenum.Reject:
		switch currentStatus {
		case statusenum.Submitted:
			newStatus = statusenum.Rejected
			activityLogMessage = fmt.Sprintf("Role was rejected by %s", payload.Modifier)
		case statusenum.ActiveEditSubmitted, statusenum.ActiveInactivationSubmitted:
			newStatus = statusenum.Active
			activityLogMessage = fmt.Sprintf("Role changes was rejected by %s", payload.Modifier)
		}
	}
	toBeUpdatedRole.Status = string(newStatus)

	/*
		Flows
		1. action == approve, currentStatus == Submitted => newStatus = Active
		2. action == approve, currentStatus == ActiveEditSubmitted => newStatus = Active
		3. action == reject, currentStatus == Submitted => newStatus = Rejected
		4. action == reject, currentStatus == ActiveEditSubmitted => newStatus => Active => copy temprole into role
	*/

	switch reviewenum.Action(payload.Action) {
	case reviewenum.Approve:
		existingResources, err := i.roleSqlAdapter.GetRoleResourceMappings(ctx, []int32{payload.ExistingRole.ID})
		if err != nil {
			return err
		}

		switch newStatus {
		case statusenum.Active:
			var payloadResources []*model.RoleResourceMapping
			if err = json.Unmarshal([]byte(*payload.ExistingRole.Resources), &payloadResources); err != nil {
				return eris.Wrap(err, "error occurred during JSON.unmarshal payloadResources")
			}

			if err = i.deleteExistingRoleResourceMappings(ctx, existingResources, payloadResources, payload.Modifier); err != nil {
				return err
			}

			if err = i.insertNewRoleResourceMappings(ctx, existingResources, payloadResources); err != nil {
				return err
			}

			policies, err := i.casbinService.GetAuthorizePolicyPayload(ctx, payloadResources)
			if err != nil {
				return err
			}

			if err = i.casbinService.StorePoliciesIntoDB(policies); err != nil {
				return err
			}
		case statusenum.Inactive:
			if err = i.roleSqlAdapter.DeleteRoleResourceMappings(ctx, sliceutil.Map(existingResources, func(dt *model.RoleResourceMapping) int32 {
				return dt.ID
			}), payload.Modifier); err != nil {
				return err
			}

			if err = i.RevokeUserRoles(ctx, payload); err != nil {
				return err
			}
		}
	case reviewenum.Reject:
		switch currentStatus {
		case statusenum.ActiveEditSubmitted:
			existingTempRole, err := i.roleSqlAdapter.GetTempRoleByRoleId(ctx, payload.ExistingRole.ID)
			if err != nil {
				return err
			}

			toBeUpdatedRole.Description = existingTempRole.Description
			toBeUpdatedRole.Name = existingTempRole.Name
			toBeUpdatedRole.Type = existingTempRole.Type
			toBeUpdatedRole.Resources = existingTempRole.Resources
		}
	}

	if err = i.roleSqlAdapter.UpdateRole(ctx, toBeUpdatedRole); err != nil {
		return err
	}

	if activityLogMessage != "" && payload.ExistingRole.ActivityLogID != nil {
		if err = i.activityLogHttpAdapter.Insert(ctx, *payload.ExistingRole.ActivityLogID, activityLogMessage, "SUCCESS"); err != nil {
			return err
		}
	}

	return nil
}

func (i *impl) deleteExistingRoleResourceMappings(
	ctx context.Context,
	existingData []*model.RoleResourceMapping,
	payload []*model.RoleResourceMapping,
	modifier string,
) error {
	payloadMapped := sliceutil.AssociateBy(payload, func(dt *model.RoleResourceMapping) string {
		return fmt.Sprintf("%d:%d", dt.RoleID, dt.ResourceID)
	})

	toBeRemovedDataIds := []int32{}
	for _, dt := range existingData {
		if _, isExist := payloadMapped[fmt.Sprintf("%d:%d", dt.RoleID, dt.ResourceID)]; !isExist {
			toBeRemovedDataIds = append(toBeRemovedDataIds, dt.ID)
		}
	}

	if len(toBeRemovedDataIds) > 0 {
		if err := i.roleSqlAdapter.DeleteRoleResourceMappings(ctx, toBeRemovedDataIds, modifier); err != nil {
			return err
		}
	}

	return nil
}

func (i *impl) insertNewRoleResourceMappings(
	ctx context.Context,
	existingData []*model.RoleResourceMapping,
	payload []*model.RoleResourceMapping,
) error {
	existingDataMapped := sliceutil.AssociateBy(existingData, func(dt *model.RoleResourceMapping) string {
		return fmt.Sprintf("%d:%d", dt.RoleID, dt.ResourceID)
	})

	toBeInsertedData := []*model.RoleResourceMapping{}
	for _, dt := range payload {
		if _, isExist := existingDataMapped[fmt.Sprintf("%d:%d", dt.RoleID, dt.ResourceID)]; !isExist {
			toBeInsertedData = append(toBeInsertedData, dt)
		}
	}

	if len(toBeInsertedData) > 0 {
		if err := i.roleSqlAdapter.InsertRoleResourceMappings(ctx, toBeInsertedData); err != nil {
			return err
		}
	}

	return nil
}

func (i *impl) RevokeUserRoles(
	ctx context.Context,
	payload *roleInterface.ServiceReviewRolePayload,
) error {
	userRoleMappings, err := i.userSqlAdapter.GetUserRoleMappings(ctx, []int32{payload.ExistingRole.ID}, nil)
	if err != nil {
		return err
	}

	if err = i.userSqlAdapter.DeleteUserRoleMappings(ctx, sliceutil.Map(userRoleMappings, func(dt *model.UserRoleMapping) int32 { return dt.ID }), payload.Modifier); err != nil {
		return err
	}

	users, err := i.userSqlAdapter.GetUsersByIds(ctx, sliceutil.Distinct(sliceutil.Map(userRoleMappings, func(dt *model.UserRoleMapping) int32 { return *dt.UserID })))
	if err != nil {
		return err
	}

	// Revoke users current login session, so they have to relogin
	for _, user := range users {
		userActiveAuthToken, err := i.jwtService.GetTokenValueByUser(ctx, user.Username)
		if err != nil {
			return err
		} else if userActiveAuthToken == nil {
			continue
		}

		if err = i.jwtService.RevokeToken(ctx, userActiveAuthToken.Username); err != nil {
			return err
		}
	}

	return nil
}
