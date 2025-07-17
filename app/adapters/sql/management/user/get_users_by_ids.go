package user

import (
	"context"

	"golang-auth-app/app/datasources/sql/gorm/model"
	"golang-auth-app/app/datasources/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) GetUsersByIds(ctx context.Context, userIds []int32) ([]*model.User, error) {
	q := query.Use(i.db.WithContext(ctx)).User
	qq := q.WithContext(ctx).Where(
		q.ID.In(userIds...),
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	)

	users, err := qq.Find()
	if err != nil {
		return nil, eris.Wrap(err, "error occurred during GetUserById")
	}

	return users, nil
}
