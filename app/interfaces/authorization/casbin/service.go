package casbin

import (
	"context"

	"golang-auth-app/app/adapters/sql/gorm/model"
	casbinDto "golang-auth-app/app/interfaces/authorization/casbin/dto"
)

type Service interface {
	AuthorizeAccess(ctx context.Context, payload *casbinDto.AuthorizePolicyPayload) (bool, error)

	GetAuthorizePolicyPayload(
		ctx context.Context,
		roleResourceMappings []*model.RoleResourceMapping,
	) ([]*casbinDto.AuthorizePolicyPayload, error)

	StorePoliciesIntoDB(policies []*casbinDto.AuthorizePolicyPayload) error

	LoadPoliciesFromDB() error

	SyncAllPolicies(ctx context.Context) error
}
