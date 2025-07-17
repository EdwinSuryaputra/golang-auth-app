package user

import (
	"context"
	"golang-auth-app/app/datasources/sql/gorm/model"
	"golang-auth-app/app/datasources/sql/gorm/query"
	"golang-auth-app/app/interfaces/errorcode"

	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

func (i *impl) GetUserByUsernameForLogin(ctx context.Context, username string) (*model.User, error) {
	q := query.Use(i.db.WithContext(ctx)).User
	qq := q.WithContext(ctx).Where(
		q.Username.Eq(username),
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	)

	user, err := qq.First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorcode.ErrCodeUserNotFound
		}

		return nil, eris.Wrap(err, "error occurred during GetUserByUsername")
	}

	return user, nil
}
