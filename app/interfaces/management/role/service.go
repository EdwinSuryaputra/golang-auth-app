package role

import (
	"context"

	"golang-auth-app/app/adapters/sql/gorm/model"

	roleDto "golang-auth-app/app/interfaces/management/role/dto"
)

type Service interface {
	GetList(ctx context.Context, payload *AdapterSqlGetPaginatedPayload) (*ServiceGetListRoleResult, error)

	GetDetail(ctx context.Context, encodedRoleId string) (*roleDto.ServiceGetDetailResult, error)

	CreateDraft(ctx context.Context, newRole *model.Role) error

	Update(ctx context.Context, payload *ServiceUpdateRolePayload) error

	Review(ctx context.Context, payload *ServiceReviewRolePayload) error
}
