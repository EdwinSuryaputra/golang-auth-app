package role

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"
	"golang-auth-app/app/common/errorcode"

	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

func (i *impl) GetRoleById(ctx context.Context, roleId int32) (*model.Role, error) {
	q := query.Use(i.db.WithContext(ctx)).Role
	qq := q.WithContext(ctx).Where(
		q.ID.Eq(roleId),
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	)

	data, err := qq.First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, eris.Wrap(errorcode.ErrCodeRoleNotFound, errorcode.ErrCodeRoleNotFound.Error())
		}

		return nil, eris.Wrap(err, "error occurred during GetRoleById")
	}

	return data, nil
}
