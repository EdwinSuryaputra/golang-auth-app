package role

import (
	"context"
	"fmt"

	reviewenum "golang-auth-app/app/common/enums/review"
	statusenum "golang-auth-app/app/common/enums/status"

	"golang-auth-app/app/common/errorcode"
	"golang-auth-app/app/interfaces/management/role"
	"golang-auth-app/app/routes/rest/middleware/authorization"
	publicfacingutil "golang-auth-app/app/utils/publicfacing"
	sliceutil "golang-auth-app/app/utils/slice"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type reviewRoleAPIPayload struct {
	Action string `json:"action"`
}

// @Summary review role
// @Tags role
// @Produce  json
// @Param body body reviewRoleAPIRequest true "update role API request body"
// @Success 200 {object} reviewRoleAPIResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /v1/role [patch]
// @Security basic
func review(
	router fiber.Router,
	authMiddleware authorization.AuthorizationMiddleware,
	roleService role.Service,
	roleSqlAdapter role.AdapterSQL,
) {
	routePath := fmt.Sprintf("%s/review/:roleId", prefix)
	requiredResources := []string{"ROLE_MANAGEMENT_REVIEW"}

	router.Patch(routePath, authMiddleware.Authorize(requiredResources), func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		modifier := string(c.Request().URI().Username())

		var payload *reviewRoleAPIPayload
		if err := c.BodyParser(&payload); err != nil {
			return errorcode.ErrCodeInvalidPayload
		}

		encodedRoleId := c.Params("roleId", "")
		roleId, err := publicfacingutil.Decode(strings.TrimSpace(encodedRoleId))
		if err != nil {
			return err
		}

		servicePayload, err := reviewRolePayloadValidation(ctx, roleSqlAdapter, roleId, payload, modifier)
		if err != nil {
			return err
		}

		err = roleService.Review(ctx, servicePayload)
		if err != nil {
			return err
		}

		resp, err := compileGetDetailAPIResponse(ctx, roleService, encodedRoleId)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(resp)
	})
}

func reviewRolePayloadValidation(
	ctx context.Context,
	roleSqlAdapter role.AdapterSQL,
	roleId int32,
	payload *reviewRoleAPIPayload,
	modifier string,
) (*role.ServiceReviewRolePayload, error) {
	if payload.Action == "" {
		return nil, errorcode.ErrCodeMissingAction
	} else if !sliceutil.Contains([]reviewenum.Action{reviewenum.Approve, reviewenum.Reject}, reviewenum.Action(payload.Action)) {
		return nil, errorcode.ErrCodeInvalidAction
	}

	existingRole, err := roleSqlAdapter.GetRoleById(ctx, roleId)
	if err != nil {
		return nil, err
	} else if !sliceutil.Contains([]statusenum.Status{statusenum.Submitted, statusenum.ActiveEditSubmitted, statusenum.ActiveInactivationSubmitted}, statusenum.Status(existingRole.Status)) {
		return nil, errorcode.ErrCodeInvalidStatus
	}

	return &role.ServiceReviewRolePayload{
		ExistingRole: existingRole,
		Action:       reviewenum.Action(payload.Action),
		Modifier:     modifier,
	}, nil
}
