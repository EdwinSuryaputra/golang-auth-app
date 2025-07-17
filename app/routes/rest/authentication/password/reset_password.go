package password

import (
	"fmt"

	authenticationInterface "golang-auth-app/app/interfaces/authentication"

	"golang-auth-app/app/interfaces/errorcode"

	"github.com/gofiber/fiber/v2"
)

type resetPasswordAPIPayload struct {
	Token       string `json:"token"`
	NewPassword string `json:"newPassword"`
}

type resetPasswordAPIResult struct {
	Result bool `json:"result"`
}

func resetPassword(
	router fiber.Router,
	authenticationService authenticationInterface.Service,
) {
	routePath := fmt.Sprintf("%s/reset", prefix)
	router.Post(routePath, func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		var payload *resetPasswordAPIPayload
		if err := c.BodyParser(&payload); err != nil {
			return errorcode.ErrCodeInvalidPayload
		}

		err := authenticationService.ResetPassword(ctx, payload.Token, payload.NewPassword)
		if err != nil {
			return err
		}

		return c.JSON(&resetPasswordAPIResult{
			Result: true,
		})
	})
}
