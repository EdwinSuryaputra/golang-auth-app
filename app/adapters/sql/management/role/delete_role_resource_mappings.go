package role

import (
	"context"
	"time"

	"golang-auth-app/app/datasources/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) DeleteRoleResourceMappings(
	ctx context.Context,
	roleResourceMappingIds []int32,
	modifier string,
) error {
	q := query.Use(i.db.WithContext(ctx)).RoleResourceMapping

	_, err := q.WithContext(ctx).Where(
		q.ID.In(roleResourceMappingIds...),
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	).Updates(map[string]any{
		"deleted_at": time.Now(),
		"deleted_by": modifier,
	})
	if err != nil {
		return eris.Wrap(err, "error occurred during DeleteRoleResourceMappings")
	}

	return nil
}
