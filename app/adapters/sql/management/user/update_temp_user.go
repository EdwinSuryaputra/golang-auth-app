package user

import (
	"context"

	"golang-auth-app/app/datasources/sql/gorm/model"
	"golang-auth-app/app/datasources/sql/gorm/query"
	convertutil "golang-auth-app/app/utils/convert"

	"github.com/rotisserie/eris"
	"gorm.io/gen/field"
)

func (i *impl) UpdateTempUser(
	ctx context.Context,
	payload *model.TempUser,
	isUndeleted bool,
) error {
	updatePayload := convertutil.StructToMap(payload)
	if isUndeleted {
		updatePayload["deleted_at"] = nil
		updatePayload["deleted_by"] = nil
	}

	q := query.Use(i.db.WithContext(ctx)).TempUser

	_, err := q.WithContext(ctx).
		Select(field.ALL).
		Omit(q.ID, q.UserID, q.CreatedAt, q.CreatedBy).
		Where(
			q.UserID.Eq(payload.UserID),
		).Unscoped().
		Updates(updatePayload)
	if err != nil {
		return eris.Wrap(err, "error occurred during UpdateTempUser")
	}

	return nil
}
