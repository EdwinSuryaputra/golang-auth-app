package user

import (
	"context"
	"fmt"

	userInterface "golang-auth-app/app/interfaces/management/user"
	userDto "golang-auth-app/app/interfaces/management/user/dto"

	"golang-auth-app/app/routes/rest/middleware/authorization"

	"github.com/gofiber/fiber/v2"
)

// @Summary Get detail User
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
	userService userInterface.Service,
) {
	routePath := fmt.Sprintf("%s/detail/:id", prefix)
	requiredResources := []string{"USER_MANAGEMENT_GET_DETAIL"}

	router.Get(routePath, authMiddleware.Authorize(requiredResources), func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		result, err := compileGetDetailAPIResponse(ctx, userService, c.Params("id", ""))
		if err != nil {
			return err
		}

		return c.JSON(result)
	})
}

func compileGetDetailAPIResponse(
	ctx context.Context,
	userService userInterface.Service,
	encodedUserId string,
) (*userDto.ServiceGetDetailResult, error) {
	data, err := userService.GetDetail(ctx, encodedUserId)
	if err != nil {
		return nil, err
	}

	return data, nil
}
