package role

import (
	"context"
	"golang-auth-app/app/datasources/sql/gorm/model"
	"golang-auth-app/app/datasources/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) GetRoleResourceMappings(ctx context.Context, roleIds []int32) ([]*model.RoleResourceMapping, error) {
	q := query.Use(i.db.WithContext(ctx)).RoleResourceMapping
	qq := q.WithContext(ctx).Where(
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	)

	if len(roleIds) > 0 {
		qq = qq.Where(q.RoleID.In(roleIds...))
	}

	data, err := qq.Find()
	if err != nil {
		return nil, eris.Wrap(err, "error occurred during GetRoleResourceMappings")
	}

	return data, nil
}
