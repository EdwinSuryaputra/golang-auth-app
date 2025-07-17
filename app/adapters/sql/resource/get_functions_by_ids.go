package resource

import (
	"context"
	"golang-auth-app/app/datasources/sql/gorm/model"
	"golang-auth-app/app/datasources/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) GetFunctionsByIds(ctx context.Context, functionIds []int32) ([]*model.Function, error) {
	q := query.Use(i.db.WithContext(ctx)).Function
	qq := q.WithContext(ctx).Where(
		q.ID.In(functionIds...),
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	)

	data, err := qq.Find()
	if err != nil {
		return nil, eris.Wrap(err, "error occurred during GetFunctionsByIds")
	}

	return data, nil
}
