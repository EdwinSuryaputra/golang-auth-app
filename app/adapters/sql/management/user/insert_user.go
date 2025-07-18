package user

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) InsertUser(ctx context.Context, payload *model.User) error {
	q := query.Use(i.db.WithContext(ctx)).User

	err := q.WithContext(ctx).Create(payload)
	if err != nil {
		return eris.Wrap(err, "error occurred during InsertUser")
	}

	return nil
}
