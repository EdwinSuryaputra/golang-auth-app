package role

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) InsertRoleResourceMappings(
	ctx context.Context,
	payload []*model.RoleResourceMapping,
) error {
	q := query.Use(i.db.WithContext(ctx)).RoleResourceMapping

	err := q.WithContext(ctx).CreateInBatches(payload, len(payload))
	if err != nil {
		return eris.Wrap(err, "error occurred during InsertRoleResourceMappings")
	}

	return nil
}
