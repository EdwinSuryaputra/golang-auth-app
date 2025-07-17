package burequestbucket

import (
	"context"
	"golang-auth-app/app/datasources/sql/gorm/model"
	"golang-auth-app/app/datasources/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) InsertBuRequestBucket(ctx context.Context, payload *model.BuRequestBucket) error {
	q := query.Use(i.db.WithContext(ctx)).BuRequestBucket

	err := q.WithContext(ctx).Create(payload)
	if err != nil {
		return eris.Wrap(err, "error occurred during InsertBuRequestBucket")
	}

	return nil
}
