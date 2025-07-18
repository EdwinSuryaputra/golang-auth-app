package user

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) InsertTempUser(ctx context.Context, payload *model.TempUser) error {
	q := query.Use(i.db.WithContext(ctx)).TempUser

	err := q.WithContext(ctx).Create(payload)
	if err != nil {
		return eris.Wrap(err, "error occurred during InsertTempRole")
	}

	return nil
}
