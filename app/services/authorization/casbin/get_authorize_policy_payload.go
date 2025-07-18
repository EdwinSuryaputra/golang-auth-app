package casbin

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"
	casbinDto "golang-auth-app/app/interfaces/authorization/casbin/dto"
	objectutil "golang-auth-app/app/utils/object"
	sliceutil "golang-auth-app/app/utils/slice"
)

func (i *impl) GetAuthorizePolicyPayload(
	ctx context.Context,
	roleResourceMappings []*model.RoleResourceMapping,
) ([]*casbinDto.AuthorizePolicyPayload, error) {
	resourceIds := map[int32]struct{}{}
	for _, rrm := range roleResourceMappings {
		if _, isExist := resourceIds[rrm.ResourceID]; !isExist {
			resourceIds[rrm.ResourceID] = struct{}{}
		}
	}
	resources, err := i.resourceSqlAdapter.GetResourcesByIds(ctx, objectutil.Keys(resourceIds))
	if err != nil {
		return nil, err
	}
	resourceAssociatedByIds := sliceutil.AssociateBy(resources, func(dt *model.Resource) int32 {
		return dt.ID
	})

	distinctedRoleIds := map[int32]struct{}{}
	for _, rrm := range roleResourceMappings {
		if _, isExist := distinctedRoleIds[rrm.RoleID]; !isExist {
			distinctedRoleIds[rrm.RoleID] = struct{}{}
		}
	}
	roles, err := i.roleSqlAdapter.GetRolesByIds(ctx, objectutil.Keys(distinctedRoleIds))
	if err != nil {
		return nil, err
	}
	roleAssociatedByIds := sliceutil.AssociateBy(roles, func(dt *model.Role) int32 { return dt.ID })

	policies := []*casbinDto.AuthorizePolicyPayload{}
	roleResourceGroupByResources := sliceutil.GroupBy(roleResourceMappings, func(dt *model.RoleResourceMapping) int32 {
		return dt.ResourceID
	})
	for resourceId, rrms := range roleResourceGroupByResources {
		policies = append(policies, &casbinDto.AuthorizePolicyPayload{
			RequiredResource: resourceAssociatedByIds[resourceId].Name,
			AuthorizedUserRoles: sliceutil.Map(rrms, func(dt *model.RoleResourceMapping) string {
				return roleAssociatedByIds[dt.RoleID].Name
			}),
		})
	}

	return policies, nil
}
