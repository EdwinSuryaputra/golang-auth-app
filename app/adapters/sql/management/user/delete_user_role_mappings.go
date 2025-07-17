package user

import (
	"context"
	"time"

	"golang-auth-app/app/datasources/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) DeleteUserRoleMappings(
	ctx context.Context,
	userRoleMappingIds []int32,
	modifier string,
) error {
	q := query.Use(i.db.WithContext(ctx)).UserRoleMapping

	for _, id := range userRoleMappingIds {
		_, err := q.WithContext(ctx).Where(
			q.ID.Eq(id),
			q.DeletedAt.IsNull(),
			q.DeletedBy.IsNull(),
		).Updates(map[string]any{
			"deleted_at": time.Now(),
			"deleted_by": modifier,
		})
		if err != nil {
			return eris.Wrap(err, "error occurred during DeleteUserRoleMappings")
		}
	}

	return nil
}
