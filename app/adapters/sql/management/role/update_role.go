package role

import (
	"context"

	"golang-auth-app/app/datasources/sql/gorm/model"
	"golang-auth-app/app/datasources/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) UpdateRole(ctx context.Context, payload *model.Role) error {
	q := query.Use(i.db.WithContext(ctx)).Role

	_, err := q.WithContext(ctx).
		Omit(q.ID, q.CreatedAt, q.CreatedBy, q.DeletedAt, q.DeletedBy).
		Where(
			q.ID.Eq(payload.ID),
			q.DeletedAt.IsNull(),
			q.DeletedBy.IsNull(),
		).
		Updates(payload)
	if err != nil {
		return eris.Wrap(err, "error occurred during UpdateRole")
	}

	return nil
}
