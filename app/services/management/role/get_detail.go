package role

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"golang-auth-app/app/datasources/sql/gorm/model"

	roleDto "golang-auth-app/app/interfaces/management/role/dto"

	objectutil "golang-auth-app/app/utils/object"
	publicfacingutil "golang-auth-app/app/utils/publicfacing"
	sliceutil "golang-auth-app/app/utils/slice"

	"github.com/rotisserie/eris"
)

func (i *impl) GetDetail(
	ctx context.Context,
	encodedRoleId string,
) (*roleDto.ServiceGetDetailResult, error) {
	roleId, err := publicfacingutil.Decode(strings.TrimSpace(encodedRoleId))
	if err != nil {
		return nil, err
	}

	existingRole, err := i.roleSqlAdapter.GetRoleById(ctx, roleId)
	if err != nil {
		return nil, err
	}

	compiledResources := []*roleDto.ServiceGetDetailResource{}
	if existingRole.Resources != nil {
		var existingRoleResources []*model.RoleResourceMapping
		err = json.Unmarshal([]byte(*existingRole.Resources), &existingRoleResources)
		if err != nil {
			return nil, eris.Wrap(err, "error occurred during JSON unmarshal resources")
		}

		if len(existingRoleResources) > 0 {
			compiledResources, err = i.getDetailCompileResources(ctx, existingRole.ApplicationID, existingRoleResources)
			if err != nil {
				return nil, err
			}
		}
	}

	return &roleDto.ServiceGetDetailResult{
		Id:            encodedRoleId,
		Name:          existingRole.Name,
		Description:   existingRole.Description,
		Type:          existingRole.Type,
		Status:        existingRole.Status,
		Resources:     compiledResources,
		InactiveDate:  existingRole.InactiveDate,
		CreatedAt:     existingRole.CreatedAt,
		CreatedBy:     existingRole.CreatedBy,
		UpdatedAt:     existingRole.UpdatedAt,
		UpdatedBy:     existingRole.UpdatedBy,
		ActivityLogId: existingRole.ActivityLogID,
	}, nil
}

func (i *impl) getDetailCompileResources(
	ctx context.Context,
	applicationId int32,
	existingRoleResources []*model.RoleResourceMapping,
) ([]*roleDto.ServiceGetDetailResource, error) {
	resourceIds := sliceutil.Map(existingRoleResources, func(dt *model.RoleResourceMapping) int32 { return dt.ResourceID })

	existingResources, err := i.resourceSqlAdapter.GetHierarchyResources(ctx, []int32{applicationId}, nil, nil, nil, resourceIds)
	if err != nil {
		return nil, err
	}
	existingApps := existingResources.Applications[applicationId]

	result := map[string]*roleDto.ServiceGetDetailResource{}
	for _, m := range existingApps.Menus {
		menuId, err := publicfacingutil.Encode(m.Id)
		if err != nil {
			return nil, err
		}

		for _, sm := range m.SubMenus {
			submenuId, err := publicfacingutil.Encode(sm.Id)
			if err != nil {
				return nil, err
			}

			key := fmt.Sprintf("%s:%s", menuId, submenuId)
			_, isMenuSubExists := result[key]
			if !isMenuSubExists {
				functions := []*roleDto.ServiceGetDetailRoleResourceFunction{}
				for _, f := range sm.Functions {
					functionId, err := publicfacingutil.Encode(f.Id)
					if err != nil {
						return nil, err
					}

					functions = append(functions, &roleDto.ServiceGetDetailRoleResourceFunction{
						Id:   functionId,
						Name: f.Name,
					})
				}

				result[key] = &roleDto.ServiceGetDetailResource{
					MenuId:      menuId,
					MenuName:    m.Name,
					SubmenuId:   submenuId,
					SubmenuName: sm.Name,
					Functions:   functions,
				}
			}
		}
	}

	return objectutil.Values(result), nil
}
