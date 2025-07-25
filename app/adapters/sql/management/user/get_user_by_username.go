package user

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"

	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

func (i *impl) GetUserByUsername(ctx context.Context, username string) (*model.User, error) {
	q := query.Use(i.db.WithContext(ctx)).User
	qq := q.WithContext(ctx).Where(
		q.Username.Eq(username),
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	)

	user, err := qq.First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, eris.Wrap(err, "error occurred during GetUserByUsername")
	}

	return user, nil
}
