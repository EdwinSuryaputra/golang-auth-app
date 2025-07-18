package resource

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) GetSubmenusByIds(ctx context.Context, submenuIds []int32) ([]*model.Submenu, error) {
	q := query.Use(i.db.WithContext(ctx)).Submenu
	qq := q.WithContext(ctx).Where(
		q.ID.In(submenuIds...),
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	)

	data, err := qq.Find()
	if err != nil {
		return nil, eris.Wrap(err, "error occurred during GetSubmenusByIds")
	}

	return data, nil
}
