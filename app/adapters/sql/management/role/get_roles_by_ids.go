package role

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) GetRolesByIds(ctx context.Context, roleIds []int32) ([]*model.Role, error) {
	q := query.Use(i.db.WithContext(ctx)).Role
	qq := q.WithContext(ctx).Where(
		q.ID.In(roleIds...),
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	)

	data, err := qq.Find()
	if err != nil {
		return nil, eris.Wrap(err, "error occurred during GetRoleById")
	}

	return data, nil
}
