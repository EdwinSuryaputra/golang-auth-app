package burequestbucket

import (
	"context"
	"golang-auth-app/app/datasources/sql/gorm/model"
	"golang-auth-app/app/datasources/sql/gorm/query"

	"github.com/rotisserie/eris"
	"gorm.io/gorm/clause"
)

func (i *impl) Upsert(ctx context.Context, payload *model.BuRequestBucket) error {
	q := query.Use(i.db.WithContext(ctx)).BuRequestBucket

	err := q.WithContext(ctx).Clauses(
		clause.OnConflict{
			Columns: []clause.Column{{Name: "id"}},
			DoUpdates: clause.AssignmentColumns([]string{"request_date",
				"user_id", "business_unit_level", "business_unit_id", "status",
				"updated_at", "updated_by"}),
			Where: clause.Where{Exprs: []clause.Expression{
				clause.Expr{SQL: "deleted_at IS NULL and deleted_by IS NULL"},
			}},
		},
	).Create(payload)
	if err != nil {
		return eris.Wrap(err, "error occurred during Upsert")
	}

	return nil
}
