package user

import (
	"context"
	"fmt"

	statusenum "golang-auth-app/app/common/enums/status"
	userenum "golang-auth-app/app/common/enums/user"
	"golang-auth-app/app/datasources/sql/gorm/model"

	userDto "golang-auth-app/app/interfaces/management/user/dto"
)

func (i *impl) Update(
	ctx context.Context,
	payload *userDto.ServiceUpdatePayload,
) error {
	var err error

	if payload.CurrentDataUser.Status == statusenum.Active.ToString() {
		existingTempUser, err := i.userSqlAdapter.GetTempUserByUserId(ctx, payload.CurrentDataUser.ID)
		if err != nil {
			return err
		}

		tempUserPayload := &model.TempUser{
			UserID:                       payload.CurrentDataUser.ID,
			Username:                     payload.CurrentDataUser.Username,
			Email:                        payload.CurrentDataUser.Email,
			FullName:                     payload.CurrentDataUser.FullName,
			PhoneNumber:                  payload.CurrentDataUser.PhoneNumber,
			Description:                  payload.CurrentDataUser.Description,
			BusinessUnitLevel:            payload.CurrentDataUser.BusinessUnitLevel,
			BusinessUnitLocationID:       payload.CurrentDataUser.BusinessUnitLocationID,
			BusinessUnitLocation:         payload.CurrentDataUser.BusinessUnitLocation,
			BusinessUnitAssignmentStatus: payload.CurrentDataUser.BusinessUnitAssignmentStatus,
			SupplierID:                   payload.CurrentDataUser.SupplierID,
			SupplierName:                 payload.CurrentDataUser.SupplierName,
			AssignedRoles:                payload.CurrentDataUser.AssignedRoles,
			CreatedBy:                    payload.CurrentDataUser.CreatedBy,
			CreatedAt:                    payload.CurrentDataUser.CreatedAt,
			UpdatedBy:                    payload.CurrentDataUser.UpdatedBy,
			UpdatedAt:                    payload.CurrentDataUser.UpdatedAt,
		}

		if existingTempUser != nil {
			if err = i.userSqlAdapter.UpdateTempUser(ctx, tempUserPayload, true); err != nil {
				return err
			}
		} else {
			if err = i.userSqlAdapter.InsertTempUser(ctx, tempUserPayload); err != nil {
				return err
			}
		}
	}

	if err = i.userSqlAdapter.UpdateUser(ctx, payload.NewDataUser); err != nil {
		return err
	}

	if userenum.UserType(payload.NewDataUser.Type) == userenum.Internal &&
		payload.NewDataUser.BusinessUnitLevel != nil &&
		payload.NewDataUser.BusinessUnitLocationID != nil &&
		payload.NewDataUser.BusinessUnitLocation != nil {
		if err = i.burbSqlAdapter.InsertBuRequestBucket(ctx, &model.BuRequestBucket{
			UserID:                 payload.NewDataUser.ID,
			BusinessUnitLevel:      *payload.NewDataUser.BusinessUnitLevel,
			BusinessUnitLocationID: *payload.NewDataUser.BusinessUnitLocationID,
			BusinessUnitLocation:   *payload.NewDataUser.BusinessUnitLocation,
			Status:                 *payload.NewDataUser.BusinessUnitAssignmentStatus,
			CreatedBy:              payload.NewDataUser.UpdatedBy,
			CreatedAt:              payload.NewDataUser.UpdatedAt,
		}); err != nil {
			return err
		}
	}

	if payload.CurrentDataUser.ActivityLogID != nil {
		var message string
		switch statusenum.Status(payload.NewDataUser.Status) {
		case statusenum.Submitted:
			message = fmt.Sprintf("User was submitted by %s", payload.NewDataUser.UpdatedBy)
		case statusenum.Draft, statusenum.ActiveEditSubmitted:
			message = fmt.Sprintf("User was edited by %s", payload.NewDataUser.UpdatedBy)
		case statusenum.ActiveInactivationSubmitted:
			message = fmt.Sprintf("User was being inactivated by %s", payload.NewDataUser.UpdatedBy)
		}

		if message != "" {
			err = i.activityLogHttpAdapter.Insert(ctx, *payload.CurrentDataUser.ActivityLogID, message, "SUCCESS")
			if err != nil {
				return err
			}
		}
	}

	return nil
}
