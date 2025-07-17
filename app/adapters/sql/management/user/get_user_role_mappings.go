package user

import (
	"context"
	"golang-auth-app/app/datasources/sql/gorm/model"
	"golang-auth-app/app/datasources/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) GetUserRoleMappings(
	ctx context.Context,
	roleIds []int32,
	userIds []int32,
) ([]*model.UserRoleMapping, error) {
	q := query.Use(i.db.WithContext(ctx)).UserRoleMapping
	qq := q.WithContext(ctx).Where(
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	)

	if len(roleIds) > 0 {
		qq = qq.Where(q.RoleID.In(roleIds...))
	}

	if len(userIds) > 0 {
		qq = qq.Where(q.UserID.In(userIds...))
	}

	data, err := qq.Find()
	if err != nil {
		return nil, eris.Wrap(err, "error occurred during GetUserRoleMappings")
	}

	return data, nil
}
