package application

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"
	"golang-auth-app/app/common/errorcode"

	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

func (i *impl) GetApplicationById(ctx context.Context, applicationId int32) (*model.Application, error) {
	q := query.Use(i.db.WithContext(ctx)).Application
	result, err := q.WithContext(ctx).Where(
		q.ID.Eq(applicationId),
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, eris.Wrap(errorcode.ErrCodeApplicationNotFound, errorcode.ErrCodeApplicationNotFound.Error())
		}

		return nil, eris.Wrap(err, "error occurred during GetApplicationById")
	}

	return result, nil
}
