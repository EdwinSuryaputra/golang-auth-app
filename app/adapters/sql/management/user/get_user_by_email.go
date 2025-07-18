package user

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"
	"golang-auth-app/app/common/errorcode"

	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

func (i *impl) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	q := query.Use(i.db.WithContext(ctx)).User
	qq := q.WithContext(ctx).Where(
		q.Email.Eq(email),
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	)

	user, err := qq.First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorcode.ErrCodeUserNotFound
		}

		return nil, eris.Wrap(err, "error occurred during GetUserById")
	}

	return user, nil
}
