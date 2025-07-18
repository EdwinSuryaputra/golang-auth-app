package role

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"
	"golang-auth-app/app/interfaces/management/role"
)

func (i *impl) GetRolesByNames(ctx context.Context, payload *role.AdapterGetRolesByNamesPayload) ([]*model.Role, error) {
	role := query.Use(i.db.WithContext(ctx)).Role

	qq := role.WithContext(ctx).Where(
		role.Name.In(payload.RoleNames...),
		role.DeletedAt.IsNull(),
		role.DeletedBy.IsNull(),
	)

	if payload.ActiveOnly {
		qq = qq.Where(role.InactiveDate.IsNull())
	}

	roles, err := qq.Find()
	if err != nil {
		return nil, err
	}

	return roles, nil
}
