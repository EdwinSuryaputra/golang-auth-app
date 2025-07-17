package resource

import (
	"context"

	resourceDto "golang-auth-app/app/interfaces/resource/dto"
	objectutil "golang-auth-app/app/utils/object"
	publicfacingutil "golang-auth-app/app/utils/publicfacing"
	sliceutil "golang-auth-app/app/utils/slice"
)

func (i *impl) GetPublicResource(
	ctx context.Context,
	payload *resourceDto.GetPublicResourcePayload,
) (*resourceDto.ResourcePubObj, error) {
	resources, err := i.resourceSqlAdapter.GetHierarchyResources(ctx, payload.ApplicationIds, nil, nil, nil, payload.ResourceIds)
	if err != nil {
		return nil, err
	}

	applications := []*resourceDto.ApplicationPubObj{}
	for _, app := range resources.Applications {
		encodedAppId, err := publicfacingutil.Encode(app.Id)
		if err != nil {
			return nil, err
		}

		menus := []*resourceDto.MenuResourcePubObj{}
		for _, menu := range app.Menus {
			encodedMenuId, err := publicfacingutil.Encode(menu.Id)
			if err != nil {
				return nil, err
			}

			submenus := []*resourceDto.SubMenuResourcePubObj{}
			for _, submenu := range menu.SubMenus {
				encodedSubmenuId, err := publicfacingutil.Encode(submenu.Id)
				if err != nil {
					return nil, err
				}

				functions, err := sliceutil.MapWithError(objectutil.Values(submenu.Functions), func(dt *resourceDto.FunctionResourceObj) (*resourceDto.FunctionResourcePubObj, error) {
					encodedFunctionId, err := publicfacingutil.Encode(dt.Id)
					if err != nil {
						return nil, err
					}

					return &resourceDto.FunctionResourcePubObj{
						Id:         encodedFunctionId,
						SubmenuId:  encodedSubmenuId,
						Name:       dt.Name,
						PublicName: dt.PublicName,
					}, nil
				})
				if err != nil {
					return nil, err
				}

				submenus = append(submenus, &resourceDto.SubMenuResourcePubObj{
					Id:         encodedSubmenuId,
					MenuId:     encodedMenuId,
					Name:       submenu.Name,
					PublicName: submenu.PublicName,
					Functions:  functions,
				})
			}

			menus = append(menus, &resourceDto.MenuResourcePubObj{
				Id:            encodedMenuId,
				ApplicationId: encodedAppId,
				Name:          menu.Name,
				PublicName:    menu.PublicName,
				SubMenus:      submenus,
			})
		}

		applications = append(applications, &resourceDto.ApplicationPubObj{
			Id:         encodedAppId,
			Name:       app.Name,
			PublicName: app.PublicName,
			Menus:      menus,
		})
	}

	return &resourceDto.ResourcePubObj{
		Applications: applications,
	}, nil
}
