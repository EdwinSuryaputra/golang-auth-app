package authorization

import (
	applicationenum "golang-auth-app/app/common/enums/application"
	casbinInterface "golang-auth-app/app/interfaces/authorization/casbin"
	casbinDto "golang-auth-app/app/interfaces/authorization/casbin/dto"
	jwtDto "golang-auth-app/app/interfaces/authorization/jwt/dto"

	jwtInterface "golang-auth-app/app/interfaces/authorization/jwt"
	"golang-auth-app/app/interfaces/errorcode"
	ctxutil "golang-auth-app/app/utils/ctx"
	sliceutil "golang-auth-app/app/utils/slice"

	"github.com/gofiber/fiber/v2"
)

type AuthorizationMiddleware interface {
	Authorize(requiredResources []string) fiber.Handler
}

type impl struct {
	jwtService    jwtInterface.Service
	casbinService casbinInterface.Service
}

func New(
	jwtService jwtInterface.Service,
	casbinService casbinInterface.Service,
) AuthorizationMiddleware {
	return &impl{
		jwtService:    jwtService,
		casbinService: casbinService,
	}
}

func (i *impl) Authorize(requiredResources []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return errorcode.ErrCodeMissingAuthToken
		}

		applicationString := c.Get("Application")
		if isValidApp := applicationenum.IsValidEnum(applicationString); !isValidApp {
			return errorcode.ErrCodeMissingApplication
		}
		c.Locals("application", applicationString)

		authTokenValue, err := i.jwtService.GetAuthToken(ctx, tokenString)
		if err != nil {
			return err
		} else if authTokenValue == nil {
			return errorcode.ErrCodeInvalidTokenExpired
		}
		c.Locals("authToken", tokenString)

		c.Request().URI().SetUsername(authTokenValue.Username)
		c.Locals("username", authTokenValue.Username)

		ctxutil.SyncLocalsToContext(c, "application", "authToken", "username")

		authorizedUserRoles := sliceutil.Map(authTokenValue.Roles, func(dt *jwtDto.AuthTokenRoleValue) string {
			return dt.Name
		})
		if sliceutil.Contains(authorizedUserRoles, "SUPERUSER") {
			return c.Next()
		}

		var isAuthorized bool
		for _, r := range requiredResources {
			isAuthorized, err = i.casbinService.AuthorizeAccess(ctx, &casbinDto.AuthorizePolicyPayload{
				RequiredResource:    r,
				AuthorizedUserRoles: authorizedUserRoles,
			})
			if err != nil {
				return err
			}

			if isAuthorized {
				return c.Next()
			}
		}

		return errorcode.ErrCodeNoPermissionAccess
	}
}
