package role

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-auth-app/config"
	"time"

	statusenum "golang-auth-app/app/common/enums/status"

	"golang-auth-app/app/datasources/sql/gorm/model"

	"golang-auth-app/app/interfaces/application"
	"golang-auth-app/app/interfaces/errorcode"
	"golang-auth-app/app/interfaces/management/role"
	"golang-auth-app/app/interfaces/resource"
	"golang-auth-app/app/routes/rest/middleware/authorization"

	publicfacingutil "golang-auth-app/app/utils/publicfacing"
	sliceutil "golang-auth-app/app/utils/slice"

	"github.com/gofiber/fiber/v2"
	"github.com/rotisserie/eris"
)

type updateRoleAPIPayload struct {
	Name        string                           `json:"name"`
	Description string                           `json:"description"`
	Type        string                           `json:"type"`
	Resources   []*updateRoleAPIPayloadResources `json:"resources"`
	IsSubmit    bool                             `json:"isSubmit"`
	IsInactive  bool                             `json:"isInactive"`
}

type updateRoleAPIPayloadResources struct {
	MenuId      string   `json:"menuId"`
	SubmenuId   string   `json:"submenuId"`
	FunctionIds []string `json:"functionIds"`
}

// @Summary update a role
// @Tags role
// @Produce  json
// @Param body body updateroleAPIRequest true "update role API request body"
// @Success 200 {object} updateroleAPIResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /v1/role [put]
// @Security basic
func update(
	router fiber.Router,
	authMiddleware authorization.AuthorizationMiddleware,
	appSqlAdapter application.AdapterSQL,
	roleSqlAdapter role.AdapterSQL,
	roleService role.Service,
	resourceSqlAdapter resource.AdapterSQL,
) {
	routePath := fmt.Sprintf("%s/:roleId", prefix)
	requiredResources := []string{"NTE_ROLE_MANAGEMENT_UPDATE"}

	router.Put(routePath, authMiddleware.Authorize(requiredResources), func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		modifier := string(c.Request().URI().Username())

		var payload *updateRoleAPIPayload
		if err := c.BodyParser(&payload); err != nil {
			return errorcode.ErrCodeInvalidPayload
		}

		encodedRoleId := c.Params("roleId", "")
		roleId, err := publicfacingutil.Decode(encodedRoleId)
		if err != nil {
			return eris.Wrap(err, "error occurred during decode public id")
		}

		servicePayload, err := updateRolePayloadValidation(ctx, roleSqlAdapter, resourceSqlAdapter, roleId, payload, modifier)
		if err != nil {
			return err
		}

		if err = roleService.Update(ctx, servicePayload); err != nil {
			return err
		}

		resp, err := compileGetDetailAPIResponse(ctx, roleService, encodedRoleId)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(resp)
	})
}

func updateRolePayloadValidation(
	ctx context.Context,
	roleSqlAdapter role.AdapterSQL,
	resourceSqlAdapter resource.AdapterSQL,
	roleId int32,
	payload *updateRoleAPIPayload,
	modifier string,
) (*role.ServiceUpdateRolePayload, error) {
	now := time.Now()

	cfg := config.Module.Role.PayloadValidation

	if payload.Name == "" {
		return nil, errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest, "Name is required")
	} else if len(payload.Name) < cfg.NameMinDigit {
		return nil, errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest, fmt.Sprintf("Name's min length is %d digit", cfg.NameMinDigit))
	}

	currentDataRole, err := roleSqlAdapter.GetRoleById(ctx, roleId)
	if err != nil {
		return nil, err
	}

	eligibleStatuses := []statusenum.Status{statusenum.Draft, statusenum.Rejected, statusenum.Active, statusenum.ActiveRejectSubmitted}
	if !sliceutil.Contains(eligibleStatuses, statusenum.Status(currentDataRole.Status)) {
		return nil, errorcode.WithCustomMessage(errorcode.ErrCodeInvalidStatus, "Invalid status")
	}

	if payload.IsInactive {
		if currentDataRole.Status != string(statusenum.Active) {
			return nil, errorcode.WithCustomMessage(errorcode.ErrCodeInvalidStatus,
				"Unable to proceed due to document status is not active yet")
		}

		return &role.ServiceUpdateRolePayload{
			CurrentDataRole: currentDataRole,
			NewDataRole: &model.Role{
				ID:           roleId,
				InactiveDate: &now,
				Status:       statusenum.ActiveInactivationSubmitted.ToString(),
				UpdatedBy:    modifier,
				UpdatedAt:    now,
			},
		}, nil
	}

	existingStatus := statusenum.Status(currentDataRole.Status)
	var newStatus statusenum.Status

	switch existingStatus {
	case statusenum.Draft:
		if payload.IsSubmit {
			newStatus = statusenum.Submitted
		}
	case statusenum.Active:
		newStatus = statusenum.ActiveEditSubmitted
	default:
		return nil, errorcode.WithCustomMessage(errorcode.ErrCodeInvalidStatus,
			fmt.Sprintf("Unable to proceed due to document status is already %s", currentDataRole.Status))
	}

	payloadRoleResourceMappings := []*model.RoleResourceMapping{}
	if len(payload.Resources) < 1 {
		return nil, errorcode.ErrCodeResourcesRequired
	} else {
		menuPayloads, err := sliceutil.MapWithError(payload.Resources, func(dt *updateRoleAPIPayloadResources) (int32, error) { return publicfacingutil.Decode(dt.MenuId) })
		if err != nil {
			return nil, err
		}

		submenuPayloads, err := sliceutil.MapWithError(payload.Resources, func(dt *updateRoleAPIPayloadResources) (int32, error) { return publicfacingutil.Decode(dt.SubmenuId) })
		if err != nil {
			return nil, err
		}

		existingResources, err := resourceSqlAdapter.GetHierarchyResources(ctx,
			[]int32{currentDataRole.ApplicationID}, menuPayloads, submenuPayloads, nil, nil)
		if err != nil {
			return nil, err
		} else if len(existingResources.Applications) < 1 {
			return nil, errorcode.WithCustomMessage(errorcode.ErrCodeNotFound, "resources not found")
		}

		for _, dt := range payload.Resources {
			menuId, _ := publicfacingutil.Decode(dt.MenuId)
			submenuId, _ := publicfacingutil.Decode(dt.SubmenuId)

			menu, isMenuExist := existingResources.Applications[currentDataRole.ApplicationID].Menus[menuId]
			if !isMenuExist {
				return nil, errorcode.WithCustomMessage(errorcode.ErrCodeInvalidMenu, fmt.Sprintf("invalid menu: %s", dt.MenuId))
			}

			submenu, isSubmenuExist := menu.SubMenus[submenuId]
			if !isSubmenuExist {
				return nil, errorcode.WithCustomMessage(errorcode.ErrCodeInvalidSubmenu, fmt.Sprintf("invalid submenu: %s", dt.SubmenuId))
			}

			if len(dt.FunctionIds) < 1 {
				return nil, errorcode.ErrCodeFunctionsRequired
			} else {
				for _, fs := range dt.FunctionIds {
					functionId, err := publicfacingutil.Decode(fs)
					if err != nil {
						return nil, err
					}

					function, isFunctionExist := submenu.Functions[functionId]
					if !isFunctionExist {
						return nil, errorcode.WithCustomMessage(errorcode.ErrCodeInvalidFunction,
							fmt.Sprintf("invalid function: %s on menu: %s, submenu: %s", fs, dt.MenuId, dt.SubmenuId))
					}

					payloadRoleResourceMappings = append(payloadRoleResourceMappings,
						&model.RoleResourceMapping{
							RoleID:     currentDataRole.ID,
							ResourceID: function.ResourceId,
							CreatedBy:  modifier,
							CreatedAt:  now,
							UpdatedBy:  modifier,
							UpdatedAt:  now,
						})
				}
			}
		}
	}

	marshalledRoleResource, err := json.Marshal(payloadRoleResourceMappings)
	if err != nil {
		return nil, eris.Wrap(err, "error occurred during stringify resources")
	}
	resources := string(marshalledRoleResource)

	return &role.ServiceUpdateRolePayload{
		CurrentDataRole: currentDataRole,
		NewDataRole: &model.Role{
			ID:            roleId,
			ApplicationID: currentDataRole.ApplicationID,
			InactiveDate:  nil,
			Name:          payload.Name,
			Description:   payload.Description,
			Type:          payload.Type,
			Status:        string(newStatus),
			Resources:     &resources,
			CreatedAt:     currentDataRole.CreatedAt,
			CreatedBy:     currentDataRole.CreatedBy,
			UpdatedBy:     modifier,
			UpdatedAt:     now,
		},
	}, nil
}
