package password

import (
	"fmt"
	authenticationInterface "golang-auth-app/app/interfaces/authentication"

	"golang-auth-app/app/interfaces/errorcode"

	"github.com/gofiber/fiber/v2"
)

type forgotPasswordAPIPayload struct {
	Email string `json:"email"`
}

type forgotPasswordAPIResult struct {
	IsEmailSent bool `json:"isEmailSent"`
}

func forgotPassword(
	router fiber.Router,
	authenticationService authenticationInterface.Service,
) {
	routePath := fmt.Sprintf("%s/forgot", prefix)
	router.Post(routePath, func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		var payload *forgotPasswordAPIPayload
		if err := c.BodyParser(&payload); err != nil {
			return errorcode.ErrCodeInvalidPayload
		}

		err := authenticationService.ForgotPassword(ctx, payload.Email)
		if err != nil {
			return err
		}

		return c.JSON(&forgotPasswordAPIResult{
			IsEmailSent: true,
		})
	})
}
