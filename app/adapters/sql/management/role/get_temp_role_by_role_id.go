package role

import (
	"context"

	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"

	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

func (i *impl) GetTempRoleByRoleId(
	ctx context.Context,
	roleId int32,
) (*model.TempRole, error) {
	q := query.Use(i.db.WithContext(ctx)).TempRole
	qq := q.WithContext(ctx).Where(
		q.RoleID.Eq(roleId),
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	)

	data, err := qq.First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, eris.Wrap(err, "error occurred during GetTempRoleByRoleId")
	}

	return data, nil
}
