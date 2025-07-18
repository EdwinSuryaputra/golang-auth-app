package authentication

import (
	"fmt"

	"golang-auth-app/app/common/errorcode"
	authenticationInterface "golang-auth-app/app/interfaces/authentication"

	"github.com/gofiber/fiber/v2"
)

type logoutAPIResponse struct {
	Message string `json:"message"`
}

func logout(
	router fiber.Router,
	authenticationService authenticationInterface.Service,
) {
	routePath := fmt.Sprintf("%s/logout", prefix)
	router.Post(routePath, func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return errorcode.ErrCodeMissingAuthToken
		}

		err := authenticationService.Logout(ctx, tokenString)
		if err != nil {
			return err
		}

		return c.JSON(&logoutAPIResponse{
			Message: "Logout successful",
		})
	})
}
