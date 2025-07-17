package user

import (
	"context"
	"golang-auth-app/app/datasources/sql/gorm/model"
	"golang-auth-app/app/datasources/sql/gorm/query"
	"golang-auth-app/app/interfaces/errorcode"

	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

func (i *impl) GetUserById(ctx context.Context, userId int32) (*model.User, error) {
	q := query.Use(i.db.WithContext(ctx)).User
	qq := q.WithContext(ctx).Where(
		q.ID.Eq(userId),
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	)

	user, err := qq.First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, eris.Wrap(errorcode.ErrCodeUserNotFound, errorcode.ErrCodeUserNotFound.Error())
		}

		return nil, eris.Wrap(err, "error occurred during GetUserById")
	}

	return user, nil
}
