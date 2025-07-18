package role

import (
	"context"
	"fmt"

	statusenum "golang-auth-app/app/common/enums/status"
	"golang-auth-app/app/adapters/sql/gorm/model"
	roleInterface "golang-auth-app/app/interfaces/management/role"
)

func (i *impl) Update(ctx context.Context, payload *roleInterface.ServiceUpdateRolePayload) error {
	var err error

	if payload.CurrentDataRole.Status == statusenum.Active.ToString() {
		existingTempRole, err := i.roleSqlAdapter.GetTempRoleByRoleId(ctx, payload.CurrentDataRole.ID)
		if err != nil {
			return err
		}

		tempRolePayload := &model.TempRole{
			RoleID:       payload.CurrentDataRole.ID,
			Name:         payload.CurrentDataRole.Name,
			Description:  payload.CurrentDataRole.Description,
			Type:         payload.CurrentDataRole.Type,
			Resources:    payload.CurrentDataRole.Resources,
			InactiveDate: payload.CurrentDataRole.InactiveDate,
			CreatedBy:    payload.CurrentDataRole.CreatedBy,
			CreatedAt:    payload.CurrentDataRole.CreatedAt,
			UpdatedBy:    payload.CurrentDataRole.UpdatedBy,
			UpdatedAt:    payload.CurrentDataRole.UpdatedAt,
		}

		if existingTempRole != nil {
			err = i.roleSqlAdapter.UpdateTempRole(ctx, tempRolePayload, true)
			if err != nil {
				return err
			}
		} else {
			err = i.roleSqlAdapter.InsertTempRole(ctx, tempRolePayload)
			if err != nil {
				return err
			}
		}
	}

	err = i.roleSqlAdapter.UpdateRole(ctx, payload.NewDataRole)
	if err != nil {
		return err
	}

	if payload.CurrentDataRole.ActivityLogID != nil {
		var message string
		switch statusenum.Status(payload.NewDataRole.Status) {
		case statusenum.Submitted:
			message = fmt.Sprintf("Role was submitted by %s", payload.NewDataRole.UpdatedBy)
		case statusenum.Draft, statusenum.ActiveEditSubmitted:
			message = fmt.Sprintf("Role was edited by %s", payload.NewDataRole.UpdatedBy)
		case statusenum.ActiveInactivationSubmitted:
			message = fmt.Sprintf("Role was being inactivated by %s", payload.NewDataRole.UpdatedBy)
		}

		if message != "" {
			err = i.activityLogHttpAdapter.Insert(ctx, *payload.CurrentDataRole.ActivityLogID, message, "SUCCESS")
			if err != nil {
				return err
			}
		}
	}

	return nil
}
