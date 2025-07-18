package role

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) InsertTempRole(ctx context.Context, newRole *model.TempRole) error {
	q := query.Use(i.db.WithContext(ctx)).TempRole

	err := q.WithContext(ctx).Create(newRole)
	if err != nil {
		return eris.Wrap(err, "error occurred during InsertTempRole")
	}

	return nil
}
