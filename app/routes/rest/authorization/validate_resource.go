package authorization

import (
	"fmt"

	casbinInterface "golang-auth-app/app/interfaces/authorization/casbin"
	casbinDto "golang-auth-app/app/interfaces/authorization/casbin/dto"
	jwtInterface "golang-auth-app/app/interfaces/authorization/jwt"
	jwtDto "golang-auth-app/app/interfaces/authorization/jwt/dto"
	sliceutil "golang-auth-app/app/utils/slice"

	"golang-auth-app/app/interfaces/errorcode"

	"github.com/gofiber/fiber/v2"
)

type validateAPIPayload struct {
	RequiredResources []string `json:"requiredResources"`
}

type validateAPIResponse struct {
	Username string `json:"username"`
	IsValid  bool   `json:"isValid"`
}

func validateResource(
	router fiber.Router,
	jwtService jwtInterface.Service,
	casbinService casbinInterface.Service,
) {
	routePath := fmt.Sprintf("%s/validate", prefix)
	router.Post(routePath, func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return errorcode.ErrCodeMissingAuthToken
		}

		var payload *validateAPIPayload
		if err := c.BodyParser(&payload); err != nil {
			return err
		}
		if err := validateResourcePayload(payload); err != nil {
			return err
		}

		authTokenValue, err := jwtService.GetAuthToken(ctx, tokenString)
		if err != nil {
			return err
		} else if authTokenValue == nil {
			return errorcode.ErrCodeInvalidTokenExpired
		}

		returnResult := &validateAPIResponse{
			Username: authTokenValue.Username,
			IsValid:  true,
		}

		authorizedUserRoles := sliceutil.Map(authTokenValue.Roles, func(dt *jwtDto.AuthTokenRoleValue) string {
			return dt.Name
		})
		if sliceutil.Contains(authorizedUserRoles, "SUPERUSER") {
			return c.JSON(returnResult)
		}

		if len(payload.RequiredResources) > 0 {
			var isAuthorized bool
			for _, r := range payload.RequiredResources {
				isAuthorized, err = casbinService.AuthorizeAccess(ctx, &casbinDto.AuthorizePolicyPayload{
					RequiredResource:    r,
					AuthorizedUserRoles: authorizedUserRoles,
				})
				if err != nil {
					return err
				}

				if isAuthorized {
					returnResult.IsValid = isAuthorized
					return c.JSON(returnResult)
				}
			}

			return errorcode.ErrCodeNoPermissionAccess
		}

		return c.JSON(returnResult)
	})
}

func validateResourcePayload(payload *validateAPIPayload) error {
	if len(payload.RequiredResources) > 0 {
		if sliceutil.Contains(payload.RequiredResources, "") {
			return errorcode.ErrCodeMissingRequiredResource
		}
	}

	return nil
}
