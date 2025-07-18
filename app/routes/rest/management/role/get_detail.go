package role

import (
	"context"
	"fmt"

	roleInterface "golang-auth-app/app/interfaces/management/role"
	roleDto "golang-auth-app/app/interfaces/management/role/dto"

	"golang-auth-app/app/routes/rest/middleware/authorization"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get detail Role
// @Tags role
// @Produce  json
// @Param body body updateroleAPIRequest true "update role API request body"
// @Success 200 {object} getDetailAPIResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /v1/management/role/detail/:id [get]
// @Security basic
func getDetail(
	router fiber.Router,
	authMiddleware authorization.AuthorizationMiddleware,
	roleService roleInterface.Service,
) {
	routePath := fmt.Sprintf("%s/detail/:roleId", prefix)
	requiredResources := []string{"ROLE_MANAGEMENT_GET_DETAIL"}

	router.Get(routePath, authMiddleware.Authorize(requiredResources), func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		resp, err := compileGetDetailAPIResponse(ctx, roleService, c.Params("roleId", ""))
		if err != nil {
			return err
		}

		return c.JSON(resp)
	})
}

func compileGetDetailAPIResponse(
	ctx context.Context,
	roleService roleInterface.Service,
	encodedRoleId string,
) (*roleDto.ServiceGetDetailResult, error) {
	data, err := roleService.GetDetail(ctx, encodedRoleId)
	if err != nil {
		return nil, err
	}

	return data, nil
}
