package role

import (
	"fmt"

	"golang-auth-app/app/interfaces/resource"
	"golang-auth-app/app/interfaces/resource/dto"
	"golang-auth-app/app/routes/rest/middleware/authorization"

	"github.com/gofiber/fiber/v2"
)

type getAllResourcesResponse struct {
	Resources *dto.ResourcePubObj `json:"resources"`
}

func getAllResources(
	router fiber.Router,
	authMiddleware authorization.AuthorizationMiddleware,
	resourceService resource.Service,
) {
	routePath := fmt.Sprintf("%s/all-resources", prefix)
	requiredResources := []string{"ROLE_MANAGEMENT_GET_LIST"}

	router.Get(routePath, authMiddleware.Authorize(requiredResources), func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		resources, err := resourceService.GetPublicResource(ctx, &dto.GetPublicResourcePayload{
			ApplicationIds: []int32{1}, // hardcoded NTE,
		})
		if err != nil {
			return err
		}

		return c.JSON(&getAllResourcesResponse{
			Resources: resources,
		})
	})
}
