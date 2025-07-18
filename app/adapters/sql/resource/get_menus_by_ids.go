package resource

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) GetMenusByIds(ctx context.Context, menuIds []int32) ([]*model.Menu, error) {
	q := query.Use(i.db.WithContext(ctx)).Menu
	qq := q.WithContext(ctx).Where(
		q.ID.In(menuIds...),
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	)

	data, err := qq.Find()
	if err != nil {
		return nil, eris.Wrap(err, "error occurred during GetMenusByIds")
	}

	return data, nil
}
