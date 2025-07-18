package role

import (
	"context"

	statusenum "golang-auth-app/app/common/enums/status"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"

	roleDto "golang-auth-app/app/interfaces/management/role/dto"

	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

func (i *impl) GetRolesByUserId(
	ctx context.Context,
	payload *roleDto.AdapterGetRolesByUserIdPayload,
) ([]*model.Role, error) {
	q := query.Use(i.db.WithContext(ctx))

	userRoleMappingTable := q.UserRoleMapping
	userRoleMappings, err := userRoleMappingTable.
		WithContext(ctx).
		Where(
			userRoleMappingTable.UserID.Eq(payload.UserId),
			userRoleMappingTable.DeletedAt.IsNull(),
			userRoleMappingTable.DeletedBy.IsNull(),
		).Find()
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, eris.Wrap(err, "error occurred during GetRolesByUserId")
	} else if len(userRoleMappings) < 1 {
		return []*model.Role{}, nil
	}

	roleIds := []int32{}
	for _, urm := range userRoleMappings {
		roleIds = append(roleIds, *urm.RoleID)
	}

	q2 := query.Use(i.db.WithContext(ctx))
	roleTable := q2.Role

	roleQuery := roleTable.WithContext(ctx).Where(
		roleTable.ID.In(roleIds...),
		roleTable.DeletedAt.IsNull(),
		roleTable.DeletedBy.IsNull(),
	)

	if payload.IsActiveOnly {
		roleQuery = roleQuery.Where(
			roleTable.Status.In(
				statusenum.Active.ToString(),
				statusenum.ActiveEditSubmitted.ToString(),
				statusenum.ActiveInactivationSubmitted.ToString(),
			),
			roleTable.InactiveDate.IsNull(),
		)
	}

	result, err := roleQuery.Find()
	if err != nil {
		return nil, eris.Wrap(err, "error occured during GetRolesByUserId")
	}

	return result, nil
}
