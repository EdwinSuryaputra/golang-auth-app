package user

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"

	"github.com/rotisserie/eris"
)

func (i *impl) InsertUserRoleMappings(
	ctx context.Context,
	payload []*model.UserRoleMapping,
) error {
	q := query.Use(i.db.WithContext(ctx)).UserRoleMapping

	err := q.WithContext(ctx).CreateInBatches(payload, len(payload))
	if err != nil {
		return eris.Wrap(err, "error occurred during InsertRoleResourceMappings")
	}

	return nil
}
