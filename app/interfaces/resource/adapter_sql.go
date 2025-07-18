package resource

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"

	resourceDto "golang-auth-app/app/interfaces/resource/dto"
)

type AdapterSQL interface {
	GetHierarchyResources(
		ctx context.Context,
		applicationIds []int32,
		menuIds []int32,
		submenuIds []int32,
		functionIds []int32,
		resourceIds []int32,
	) (*resourceDto.ResourceObj, error)

	GetMenusByIds(ctx context.Context, menuIds []int32) ([]*model.Menu, error)

	GetSubmenusByIds(ctx context.Context, submenuIds []int32) ([]*model.Submenu, error)

	GetFunctionsByIds(ctx context.Context, functionIds []int32) ([]*model.Function, error)

	GetResourcesByIds(ctx context.Context, resourceIds []int32) ([]*model.Resource, error)
}
