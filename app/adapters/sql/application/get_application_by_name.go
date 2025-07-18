package application

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"
	"golang-auth-app/app/common/errorcode"

	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

func (i *impl) GetApplicationByName(ctx context.Context, applicationName string) (*model.Application, error) {
	q := query.Use(i.db.WithContext(ctx)).Application
	result, err := q.WithContext(ctx).Where(
		q.Name.Eq(applicationName),
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	).First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, eris.Wrap(errorcode.ErrCodeApplicationNotFound, errorcode.ErrCodeApplicationNotFound.Error())
		}

		return nil, eris.Wrap(err, "error occurred during GetApplicationByName")
	}

	return result, nil
}
