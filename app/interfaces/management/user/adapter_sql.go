package user

import (
	"context"

	"golang-auth-app/app/datasources/sql/gorm/model"

	userDto "golang-auth-app/app/interfaces/management/user/dto"
)

type AdapterSQL interface {
	GetPaginated(ctx context.Context, payload *userDto.AdapterSqlGetPaginatedPayload) (*userDto.AdapterSqlGetPaginatedResult, error)

	GetUserByUsernameForLogin(ctx context.Context, username string) (*model.User, error)
	GetUserByUsername(ctx context.Context, username string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserById(ctx context.Context, userId int32) (*model.User, error)
	GetUsersByIds(ctx context.Context, userIds []int32) ([]*model.User, error)
	InsertUser(ctx context.Context, payload *model.User) error
	UpdateUser(ctx context.Context, payload *model.User) error

	GetTempUserByUserId(ctx context.Context, userId int32) (*model.TempUser, error)
	InsertTempUser(ctx context.Context, payload *model.TempUser) error
	UpdateTempUser(ctx context.Context, payload *model.TempUser, isUndeleted bool) error

	GetUserRoleMappings(ctx context.Context, roleIds []int32, userIds []int32) ([]*model.UserRoleMapping, error)
	InsertUserRoleMappings(ctx context.Context, payload []*model.UserRoleMapping) error
	DeleteUserRoleMappings(ctx context.Context, userRoleMappingIds []int32, modifier string) error
}
