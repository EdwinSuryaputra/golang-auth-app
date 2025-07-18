package user

import (
	"fmt"

	statusenum "golang-auth-app/app/common/enums/status"

	"golang-auth-app/app/common/errorcode"
	"golang-auth-app/app/interfaces/generic"
	"golang-auth-app/app/interfaces/management/role"
	"golang-auth-app/app/interfaces/management/user"
	"golang-auth-app/app/routes/rest/middleware/authorization"
	publicfacingutil "golang-auth-app/app/utils/publicfacing"

	"github.com/gofiber/fiber/v2"
)

type deleteUserAPIResponse struct {
	Message string `json:"message"`
}

// @Summary delete a user
// @Tags user
// @Produce  json
// @Param body body deleteUserAPIRequest true "delete user API request body"
// @Success 200 {object} deleteUserAPIResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /v1/management/user [delete]
// @Security basic
func delete(
	router fiber.Router,
	authMiddleware authorization.AuthorizationMiddleware,
	genericSqlAdapter generic.AdapterSQL,
	userSqlAdapter user.AdapterSQL,
	roleSqlAdapter role.AdapterSQL,
) {
	routePath := fmt.Sprintf("%s/:id", prefix)
	requiredResources := []string{"USER_MANAGEMENT_DELETE_DRAFT"}

	router.Delete(routePath, authMiddleware.Authorize(requiredResources), func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		userId, err := publicfacingutil.Decode(c.Params("id", ""))
		if err != nil {
			return err
		}

		modifier := string(c.Request().URI().Username())

		existingUser, err := userSqlAdapter.GetUserById(ctx, userId)
		if err != nil {
			return err
		} else if existingUser.Status != string(statusenum.Draft) {
			return errorcode.WithCustomMessage(errorcode.ErrCodeUnprocessableEntity,
				"Unable to delete due to user is no longer on Draft")
		}

		err = genericSqlAdapter.SoftDelete(ctx, map[string]interface{}{
			"id":         userId,
			"deleted_by": nil,
			"deleted_at": nil,
		}, modifier)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(&deleteUserAPIResponse{
			Message: "User has been succesfully deleted",
		})
	})
}
