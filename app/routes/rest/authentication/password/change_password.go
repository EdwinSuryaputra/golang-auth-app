package password

import (
	"fmt"

	authenticationInterface "golang-auth-app/app/interfaces/authentication"

	"golang-auth-app/app/common/errorcode"

	"github.com/gofiber/fiber/v2"
)

type changePasswordAPIPayload struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type changePasswordAPIResult struct {
	Result bool `json:"result"`
}

func changePassword(
	router fiber.Router,
	authenticationService authenticationInterface.Service,
) {
	routePath := fmt.Sprintf("%s/change", prefix)
	router.Post(routePath, func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return errorcode.ErrCodeMissingAuthToken
		}

		var payload *changePasswordAPIPayload
		if err := c.BodyParser(&payload); err != nil {
			return errorcode.ErrCodeInvalidPayload
		}

		err := authenticationService.ChangePassword(ctx, tokenString, payload.OldPassword, payload.NewPassword)
		if err != nil {
			return err
		}

		return c.JSON(&changePasswordAPIResult{
			Result: true,
		})
	})
}
