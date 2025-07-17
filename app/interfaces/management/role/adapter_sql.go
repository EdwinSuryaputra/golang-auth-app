package role

import (
	"context"

	"golang-auth-app/app/datasources/sql/gorm/model"

	roleDto "golang-auth-app/app/interfaces/management/role/dto"
)

type AdapterSQL interface {
	GetPaginated(ctx context.Context, payload *AdapterSqlGetPaginatedPayload) (*AdapterSqlGetPaginatedResult, error)

	GetRolesByNames(ctx context.Context, payload *AdapterGetRolesByNamesPayload) ([]*model.Role, error)
	GetRolesByUserId(ctx context.Context, payload *roleDto.AdapterGetRolesByUserIdPayload) ([]*model.Role, error)
	GetRoleById(ctx context.Context, roleId int32) (*model.Role, error)
	GetRolesByIds(ctx context.Context, roleIds []int32) ([]*model.Role, error)
	InsertRole(ctx context.Context, newRole *model.Role) error
	UpdateRole(ctx context.Context, payload *model.Role) error

	GetRoleResourceMappings(ctx context.Context, roleIds []int32) ([]*model.RoleResourceMapping, error)
	InsertRoleResourceMappings(ctx context.Context, payload []*model.RoleResourceMapping) error
	DeleteRoleResourceMappings(ctx context.Context, roleResourceMappingIds []int32, modifier string) error

	GetTempRoleByRoleId(ctx context.Context, roleId int32) (*model.TempRole, error)

	InsertTempRole(ctx context.Context, newRole *model.TempRole) error
	UpdateTempRole(ctx context.Context, payload *model.TempRole, isUndeleted bool) error
}
