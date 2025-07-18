package password

import (
	"fmt"

	authenticationInterface "golang-auth-app/app/interfaces/authentication"

	"golang-auth-app/app/common/errorcode"

	"github.com/gofiber/fiber/v2"
)

type validateTokenAPIPayload struct {
	Token string `json:"token"`
}

type validateTokenAPIResult struct {
	Result bool `json:"result"`
}

func validateResetPasswordToken(
	router fiber.Router,
	authenticationService authenticationInterface.Service,
) {
	routePath := fmt.Sprintf("%s/reset/token/validate", prefix)
	router.Post(routePath, func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		var payload *validateTokenAPIPayload
		if err := c.BodyParser(&payload); err != nil {
			return errorcode.ErrCodeInvalidPayload
		}

		redisObj, err := authenticationService.ValidateResetPasswordToken(ctx, payload.Token)
		if err != nil {
			return err
		}

		var result bool
		if redisObj != nil && redisObj.TokenString == payload.Token {
			result = true
		}

		return c.JSON(&validateTokenAPIResult{
			Result: result,
		})
	})
}
