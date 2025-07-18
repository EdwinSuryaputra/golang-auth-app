package resource

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) GetResourcesByIds(ctx context.Context, resourceIds []int32) ([]*model.Resource, error) {
	q := query.Use(i.db.WithContext(ctx)).Resource
	qq := q.WithContext(ctx).Where(
		q.ID.In(resourceIds...),
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	)

	data, err := qq.Find()
	if err != nil {
		return nil, eris.Wrap(err, "error occurred during GetFunctionsByIds")
	}

	return data, nil
}
