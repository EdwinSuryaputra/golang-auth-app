package role

import (
	"context"
	"golang-auth-app/app/datasources/sql/gorm/model"
	"golang-auth-app/app/datasources/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) InsertRole(ctx context.Context, newRole *model.Role) error {
	q := query.Use(i.db.WithContext(ctx)).Role

	err := q.WithContext(ctx).Create(newRole)
	if err != nil {
		return eris.Wrap(err, "error occurred during InsertRole")
	}

	return nil
}
