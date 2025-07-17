package role

import (
	"context"

	"golang-auth-app/app/datasources/sql/gorm/model"
	"golang-auth-app/app/datasources/sql/gorm/query"
	convertutil "golang-auth-app/app/utils/convert"

	"github.com/rotisserie/eris"
)

func (i *impl) UpdateTempRole(ctx context.Context, payload *model.TempRole, isUndeleted bool) error {
	updatePayload := convertutil.StructToMap(payload)
	if isUndeleted {
		updatePayload["deleted_at"] = nil
		updatePayload["deleted_by"] = nil
	}

	q := query.Use(i.db.WithContext(ctx)).TempRole

	_, err := q.WithContext(ctx).
		Omit(q.ID, q.RoleID, q.CreatedAt, q.CreatedBy).
		Where(
			q.RoleID.Eq(payload.RoleID),
		).Unscoped().
		Updates(updatePayload)
	if err != nil {
		return eris.Wrap(err, "error occurred during UpdateTempRole")
	}

	return nil
}
