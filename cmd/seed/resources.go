package main

import (
	"context"
	"time"

	"golang-auth-app/app/adapters/sql/gorm/model"

	"gorm.io/gorm/clause"
)

type application struct {
	Name        string  `json:"name"`
	PublicName  string  `json:"publicName"`
	Description string  `json:"description"`
	Menus       []*menu `json:"menus"`
}

type menu struct {
	Name        string     `json:"name"`
	PublicName  string     `json:"publicName"`
	Description string     `json:"description"`
	Submenus    []*submenu `json:"submenus"`
}

type submenu struct {
	Name        string      `json:"name"`
	PublicName  string      `json:"public_name"`
	Description string      `json:"description"`
	Functions   []*function `json:"functions"`
}

type function struct {
	Name         string `json:"name"`
	PublicName   string `json:"publicName"`
	Description  string `json:"description"`
	ResourceName string `json:"resourceName"`
}

func seedResource(
	ctx context.Context,
	now time.Time,
	creator string,
) error {
	var err error

	var applicationId, menuId, submenuId,
		functionId, resourceId int32 = 1, 1, 1, 1, 1

	for _, app := range getResourceData() {
		appModel := &model.Application{
			ID:         applicationId,
			Name:       app.Name,
			PublicName: app.PublicName,
			CreatedBy:  creator,
			CreatedAt:  now,
			UpdatedBy:  creator,
			UpdatedAt:  now,
		}

		if err = upsertBatches(ctx, upsertPayload[model.Application]{
			Data:              []*model.Application{appModel},
			OnConflictColumns: []clause.Column{{Name: "id"}},
			DoUpdateColumns:   []string{"name", "public_name", "description", "updated_by", "updated_at"},
		}); err != nil {
			return err
		}

		for _, menu := range app.Menus {
			menuModel := &model.Menu{
				ID:            menuId,
				ApplicationID: appModel.ID,
				Name:          menu.Name,
				PublicName:    menu.PublicName,
				Description:   menu.Description,
				CreatedBy:     creator,
				CreatedAt:     now,
				UpdatedBy:     creator,
				UpdatedAt:     now,
			}

			if err = upsertBatches(ctx, upsertPayload[model.Menu]{
				Data:              []*model.Menu{menuModel},
				OnConflictColumns: []clause.Column{{Name: "id"}},
				DoUpdateColumns:   []string{"name", "public_name", "description", "updated_by", "updated_at"},
			}); err != nil {
				return err
			}

			for _, submenu := range menu.Submenus {
				submenuModel := &model.Submenu{
					ID:          submenuId,
					MenuID:      menuModel.ID,
					Name:        submenu.Name,
					PublicName:  submenu.PublicName,
					Description: submenu.Description,
					CreatedBy:   creator,
					CreatedAt:   now,
					UpdatedBy:   creator,
					UpdatedAt:   now,
				}

				if err = upsertBatches(ctx, upsertPayload[model.Submenu]{
					Data:              []*model.Submenu{submenuModel},
					OnConflictColumns: []clause.Column{{Name: "id"}},
					DoUpdateColumns:   []string{"name", "public_name", "description", "updated_by", "updated_at"},
				}); err != nil {
					return err
				}

				for _, function := range submenu.Functions {
					functionModel := &model.Function{
						ID:          functionId,
						SubmenuID:   submenuModel.ID,
						Name:        function.Name,
						PublicName:  function.PublicName,
						Description: function.Description,
						CreatedBy:   creator,
						CreatedAt:   now,
						UpdatedBy:   creator,
						UpdatedAt:   now,
					}

					if err = upsertBatches(ctx, upsertPayload[model.Function]{
						Data:              []*model.Function{functionModel},
						OnConflictColumns: []clause.Column{{Name: "id"}},
						DoUpdateColumns:   []string{"name", "public_name", "description", "updated_by", "updated_at"},
					}); err != nil {
						return err
					}

					resourceModel := &model.Resource{
						ID:         resourceId,
						Name:       function.ResourceName,
						MenuID:     menuModel.ID,
						SubmenuID:  submenuModel.ID,
						FunctionID: functionModel.ID,
						CreatedBy:  creator,
						CreatedAt:  now,
						UpdatedBy:  creator,
						UpdatedAt:  now,
					}

					if err = upsertBatches(ctx, upsertPayload[model.Resource]{
						Data:              []*model.Resource{resourceModel},
						OnConflictColumns: []clause.Column{{Name: "id"}},
						DoUpdateColumns:   []string{"name", "menu_id", "submenu_id", "function_id", "updated_by", "updated_at"},
					}); err != nil {
						return err
					}

					functionId++
					resourceId++
				}

				submenuId++
			}

			menuId++
		}

		applicationId++
	}

	return nil
}

func getResourceData() []*application {
	return []*application{
		{
			Name:        "A",
			PublicName:  "A",
			Description: "Application A",
			Menus: []*menu{
				{
					Name:       "SETUP",
					PublicName: "Setup",
					Submenus: []*submenu{
						{
							Name:       "ROLE_MANAGEMENT",
							PublicName: "Role Management",
							Functions: []*function{
								{
									Name:         "VIEW_LIST",
									PublicName:   "View List",
									ResourceName: "ROLE_MANAGEMENT_GET_LIST",
								},
								{
									Name:         "VIEW_DETAIL",
									PublicName:   "View Detail",
									ResourceName: "ROLE_MANAGEMENT_GET_DETAIL",
								},
								{
									Name:         "VIEW_LOG_DETAIL",
									PublicName:   "View Log Detail",
									ResourceName: "ROLE_MANAGEMENT_GET_LOG_DETAIL",
								},
								{
									Name:         "CREATE_DRAFT",
									PublicName:   "Create Draft",
									ResourceName: "ROLE_MANAGEMENT_CREATE_DRAFT",
								},
								{
									Name:         "UPDATE",
									PublicName:   "Update",
									ResourceName: "ROLE_MANAGEMENT_UPDATE",
								},
								{
									Name:         "REVIEW",
									PublicName:   "Review",
									ResourceName: "ROLE_MANAGEMENT_REVIEW",
								},
							},
						},
						{
							Name:       "USER_MANAGEMENT",
							PublicName: "User Management",
							Functions: []*function{
								{
									Name:         "VIEW_LIST",
									PublicName:   "View List",
									ResourceName: "USER_MANAGEMENT_GET_LIST",
								},
								{
									Name:         "VIEW_DETAIL",
									PublicName:   "View Detail",
									ResourceName: "USER_MANAGEMENT_GET_DETAIL",
								},
								{
									Name:         "VIEW_LOG_DETAIL",
									PublicName:   "View Log Detail",
									ResourceName: "USER_MANAGEMENT_GET_LOG_DETAIL",
								},
								{
									Name:         "CREATE_DRAFT",
									PublicName:   "Create Draft",
									ResourceName: "USER_MANAGEMENT_CREATE_DRAFT",
								},
								{
									Name:         "UPDATE",
									PublicName:   "Update",
									ResourceName: "USER_MANAGEMENT_UPDATE",
								},
								{
									Name:         "DRAFT_APPROVAL",
									PublicName:   "Draft Approval",
									ResourceName: "USER_MANAGEMENT_REVIEW",
								},
							},
						},
					},
				},
			},
		},
	}
}
