package resource

import (
	"context"

	"golang-auth-app/app/datasources/sql/gorm/query"
	resourceDto "golang-auth-app/app/interfaces/resource/dto"

	"github.com/rotisserie/eris"
)

func (i *impl) GetHierarchyResources(
	ctx context.Context,
	applicationIds []int32,
	menuIds []int32,
	submenuIds []int32,
	functionIds []int32,
	resourceIds []int32,
) (*resourceDto.ResourceObj, error) {
	q := query.Use(i.db.WithContext(ctx))

	applicationTable := q.Application
	menuTable := q.Menu
	submenuTable := q.Submenu
	functionTable := q.Function
	resourceTable := q.Resource

	qq := menuTable.WithContext(ctx).
		Join(applicationTable, applicationTable.ID.EqCol(menuTable.ApplicationID)).
		Join(submenuTable, submenuTable.MenuID.EqCol(menuTable.ID)).
		Join(functionTable, functionTable.SubmenuID.EqCol(submenuTable.ID)).
		Join(resourceTable,
			resourceTable.MenuID.EqCol(menuTable.ID),
			resourceTable.SubmenuID.EqCol(submenuTable.ID),
			resourceTable.FunctionID.EqCol(functionTable.ID),
		).
		Where(
			applicationTable.DeletedAt.IsNull(),
			applicationTable.DeletedBy.IsNull(),
			menuTable.DeletedAt.IsNull(),
			menuTable.DeletedBy.IsNull(),
			submenuTable.DeletedAt.IsNull(),
			submenuTable.DeletedBy.IsNull(),
			functionTable.DeletedAt.IsNull(),
			functionTable.DeletedBy.IsNull(),
			resourceTable.DeletedAt.IsNull(),
			resourceTable.DeletedBy.IsNull(),
		)

	if len(applicationIds) > 0 {
		qq = qq.Where(applicationTable.ID.In(applicationIds...))
	}

	if len(menuIds) > 0 {
		qq = qq.Where(menuTable.ID.In(menuIds...))
	}

	if len(submenuIds) > 0 {
		qq = qq.Where(submenuTable.ID.In(submenuIds...))
	}

	if len(functionIds) > 0 {
		qq = qq.Where(functionTable.ID.In(functionIds...))
	}

	if len(resourceIds) > 0 {
		qq = qq.Where(resourceTable.ID.In(resourceIds...))
	}

	qq = qq.Select(
		applicationTable.ID.As("ApplicationId"),
		applicationTable.Name.As("ApplicationName"),
		applicationTable.PublicName.As("ApplicationPublicName"),
		menuTable.ID.As("MenuId"),
		menuTable.Name.As("MenuName"),
		menuTable.PublicName.As("MenuPublicName"),
		submenuTable.ID.As("SubmenuId"),
		submenuTable.Name.As("SubmenuName"),
		submenuTable.PublicName.As("SubmenuPublicName"),
		functionTable.ID.As("FunctionId"),
		functionTable.Name.As("FunctionName"),
		functionTable.PublicName.As("FunctionPublicName"),
		resourceTable.ID.As("ResourceId"),
		resourceTable.Name.As("ResourceName"),
	)

	type joinedResourcesQueryResult struct {
		ApplicationId         int32
		ApplicationName       string
		ApplicationPublicName string
		MenuId                int32
		MenuName              string
		MenuPublicName        string
		SubmenuId             int32
		SubmenuName           string
		SubmenuPublicName     string
		FunctionId            int32
		FunctionName          string
		FunctionPublicName    string
		ResourceId            int32
		ResourceName          string
	}
	var queryResult []joinedResourcesQueryResult

	err := qq.Scan(&queryResult)
	if err != nil {
		return nil, eris.Wrap(err, "error occurred during GetResources")
	}

	result := &resourceDto.ResourceObj{
		Applications: make(map[int32]*resourceDto.ApplicationResourceObj),
	}

	for _, res := range queryResult {
		application, isAppExist := result.Applications[res.ApplicationId]
		if !isAppExist {
			application = &resourceDto.ApplicationResourceObj{
				Id:         res.ApplicationId,
				Name:       res.ApplicationName,
				PublicName: res.ApplicationPublicName,
				Menus:      make(map[int32]*resourceDto.MenuResourceObj),
			}
			result.Applications[res.ApplicationId] = application
		}

		menu, isMenuExist := application.Menus[res.MenuId]
		if !isMenuExist {
			menu = &resourceDto.MenuResourceObj{
				Id:            res.MenuId,
				ApplicationId: application.Id,
				Name:          res.MenuName,
				PublicName:    res.MenuPublicName,
				SubMenus:      make(map[int32]*resourceDto.SubMenuResourceObj),
			}
			application.Menus[res.MenuId] = menu
		}

		submenu, isSubmenuExist := menu.SubMenus[res.SubmenuId]
		if !isSubmenuExist {
			submenu = &resourceDto.SubMenuResourceObj{
				Id:         res.SubmenuId,
				MenuId:     menu.Id,
				Name:       res.SubmenuName,
				PublicName: res.SubmenuPublicName,
				Functions:  make(map[int32]*resourceDto.FunctionResourceObj),
			}
			menu.SubMenus[res.SubmenuId] = submenu
		}

		if _, functionExists := submenu.Functions[res.FunctionId]; !functionExists {
			submenu.Functions[res.FunctionId] = &resourceDto.FunctionResourceObj{
				Id:           res.FunctionId,
				SubMenuId:    res.SubmenuId,
				Name:         res.FunctionName,
				PublicName:   res.FunctionPublicName,
				ResourceId:   res.ResourceId,
				ResourceName: res.ResourceName,
			}
		}
	}

	return result, nil
}
