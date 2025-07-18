package user

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) UpdateUser(ctx context.Context, payload *model.User) error {
	q := query.Use(i.db.WithContext(ctx)).User

	_, err := q.WithContext(ctx).
		Omit(q.ID, q.CreatedAt, q.CreatedBy, q.DeletedAt, q.DeletedBy).
		Where(
			q.ID.Eq(payload.ID),
			q.DeletedAt.IsNull(),
			q.DeletedBy.IsNull(),
		).
		Updates(payload)
	if err != nil {
		return eris.Wrap(err, "error occurred during UpdateUser")
	}

	return nil
}
