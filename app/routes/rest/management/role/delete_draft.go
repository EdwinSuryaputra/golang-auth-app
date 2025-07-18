package role

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/rotisserie/eris"

	"golang-auth-app/app/common/errorcode"
	"golang-auth-app/app/interfaces/generic"
	"golang-auth-app/app/interfaces/management/role"
	"golang-auth-app/app/routes/rest/middleware/authorization"
	publicfacingutil "golang-auth-app/app/utils/publicfacing"
)

type deleteDraftRoleAPIResponse struct {
	Message string `json:"message"`
}

// @Summary Delete draft a Role
// @Tags role
// @Produce  json
// @Success 200 {object} deleteDraftRoleAPIResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /v1/role [deleteDraft]
// @Security basic
func deleteDraft(
	router fiber.Router,
	authMiddleware authorization.AuthorizationMiddleware,
	genericSqlAdapter generic.AdapterSQL,
	roleSqlAdapter role.AdapterSQL,
) {
	routePath := fmt.Sprintf("%s/:roleId", prefix)
	requiredResources := []string{"ROLE_MANAGEMENT_DELETE_DRAFT"}

	router.Delete(routePath, authMiddleware.Authorize(requiredResources), func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		roleId := strings.TrimSpace(c.Params("roleId"))
		if roleId == "" {
			return errorcode.ErrCodeInvalidId
		}
		decodedRoleId, err := publicfacingutil.Decode(roleId)
		if err != nil {
			return eris.Wrap(err, "error occurred during decode public id")
		}

		// modifier := string(c.Request().URI().Username())

		if existingrole, err := roleSqlAdapter.GetRoleById(ctx, decodedRoleId); err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Error occurred during get existing role")
		} else if existingrole == nil {
			return fiber.NewError(fiber.StatusUnprocessableEntity, "Role is not exist")
		}

		// if err := genericSqlAdapter.SoftDelete(ctx, map[string]interface{}{
		// 	"id":   decodedRoleId,
		// 	"deleted_by": nil,
		// 	"deleted_at": nil,
		// }, modifier); err != nil {
		// 	return fiber.NewError(fiber.StatusInternalServerError, "Error occurred during delete role")
		// }

		return c.Status(fiber.StatusOK).JSON(&deleteDraftRoleAPIResponse{
			Message: "Role has been deleted",
		})
	})
}
