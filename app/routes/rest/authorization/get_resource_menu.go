package authorization

import (
	"fmt"

	jwtInterface "golang-auth-app/app/interfaces/authorization/jwt"
	resourceDto "golang-auth-app/app/interfaces/resource/dto"

	"golang-auth-app/app/interfaces/errorcode"

	"github.com/gofiber/fiber/v2"
)

type getResourceMenuAPIResponse struct {
	Resources *resourceDto.ResourcePubObj `json:"resources"`
}

func getResourceMenu(
	router fiber.Router,
	jwtService jwtInterface.Service,
) {
	routePath := fmt.Sprintf("%s/resource/menu", prefix)
	router.Get(routePath, func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return errorcode.ErrCodeMissingAuthToken
		}

		authTokenValue, err := jwtService.GetAuthToken(ctx, tokenString)
		if err != nil {
			return err
		} else if authTokenValue == nil {
			return errorcode.ErrCodeInvalidTokenExpired
		}

		return c.JSON(&getResourceMenuAPIResponse{
			Resources: authTokenValue.Resources,
		})
	})
}
