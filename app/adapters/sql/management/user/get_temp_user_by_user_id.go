package user

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"

	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

func (i *impl) GetTempUserByUserId(ctx context.Context, userId int32) (*model.TempUser, error) {
	q := query.Use(i.db.WithContext(ctx)).TempUser
	qq := q.WithContext(ctx).Where(
		q.UserID.Eq(userId),
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	)

	data, err := qq.First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, eris.Wrap(err, "error occurred during GetTempUserByUserId")
	}

	return data, nil
}
