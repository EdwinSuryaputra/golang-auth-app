package role

import (
	"fmt"
	"golang-auth-app/config"
	"time"

	roleenum "golang-auth-app/app/common/enums/role"
	statusenum "golang-auth-app/app/common/enums/status"

	"golang-auth-app/app/datasources/sql/gorm/model"
	"golang-auth-app/app/interfaces/application"
	"golang-auth-app/app/interfaces/errorcode"
	"golang-auth-app/app/interfaces/management/role"
	"golang-auth-app/app/routes/rest/middleware/authorization"
	publicfacingutil "golang-auth-app/app/utils/publicfacing"

	"github.com/gofiber/fiber/v2"
)

type createDraftRoleAPIPayload struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
}

// @Summary Create Draft a Role
// @Tags role
// @Produce json
// @Param body body createRoleAPIRequest true "Create draft role API payload"
// @Success 200 {object} createDraftRoleAPIResponse
// @Failure 400 {object}
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} type
// @Router /v1/management/role [post]
// @Security basic
func createDraft(
	router fiber.Router,
	authMiddleware authorization.AuthorizationMiddleware,
	roleService role.Service,
	appSqlAdapter application.AdapterSQL,
	roleSqlAdapter role.AdapterSQL,
) {
	routePath := fmt.Sprintf("%s/", prefix)
	requiredResources := []string{"NTE_ROLE_MANAGEMENT_CREATE_DRAFT"}

	router.Post(routePath, authMiddleware.Authorize(requiredResources), func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		var payload createDraftRoleAPIPayload
		if err := c.BodyParser(&payload); err != nil {
			return errorcode.ErrCodeInvalidPayload
		} else if err := createRolePayloadValidation(&payload); err != nil {
			return err
		}

		modifier := string(c.Request().URI().Username())
		createdAt := time.Now()
		applicationName := c.Locals("application").(string)

		existingApp, err := appSqlAdapter.GetApplicationByName(ctx, applicationName)
		if err != nil {
			return err
		}

		if existingrole, err := roleSqlAdapter.GetRolesByNames(ctx, &role.AdapterGetRolesByNamesPayload{
			RoleNames:  []string{payload.Name},
			ActiveOnly: false,
		}); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred during get existing role")
		} else if len(existingrole) > 0 {
			return fiber.NewError(fiber.StatusUnprocessableEntity, "Role is already exist")
		}

		newRole := &model.Role{
			ApplicationID: existingApp.ID,
			Name:          payload.Name,
			Description:   payload.Description,
			Type:          payload.Type,
			Status:        string(statusenum.Draft),
			CreatedBy:     modifier,
			CreatedAt:     createdAt,
			UpdatedBy:     modifier,
			UpdatedAt:     createdAt,
		}
		if err := roleService.CreateDraft(ctx, newRole); err != nil {
			return err
		}

		encodedRoleId, err := publicfacingutil.Encode(newRole.ID)
		if err != nil {
			return err
		}

		resp, err := compileGetDetailAPIResponse(ctx, roleService, encodedRoleId)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(resp)
	})
}

func createRolePayloadValidation(payload *createDraftRoleAPIPayload) error {
	cfg := config.Module.Role.PayloadValidation

	if payload.Name == "" {
		return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest, "Name is required")
	} else if len(payload.Name) < cfg.NameMinDigit {
		return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest, fmt.Sprintf("Name's min length is %d digit", cfg.NameMinDigit))
	}

	if payload.Description == "" {
		return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest, "Description is required")
	} else if len(payload.Description) < cfg.DescriptionMinDigit {
		return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest, fmt.Sprintf("Description's min length is %d digit", cfg.DescriptionMinDigit))
	}

	if payload.Type == "" {
		return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest, "Type is required")
	} else if !roleenum.IsValidRoleType(payload.Type) {
		return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest, "Invalid type")
	}

	return nil
}
