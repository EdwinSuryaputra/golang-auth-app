package user

import (
	"context"
	"fmt"

	reviewenum "golang-auth-app/app/common/enums/review"
	statusenum "golang-auth-app/app/common/enums/status"

	"golang-auth-app/app/common/errorcode"

	userInterface "golang-auth-app/app/interfaces/management/user"
	userDto "golang-auth-app/app/interfaces/management/user/dto"

	"golang-auth-app/app/routes/rest/middleware/authorization"

	publicfacingutil "golang-auth-app/app/utils/publicfacing"
	sliceutil "golang-auth-app/app/utils/slice"

	"github.com/gofiber/fiber/v2"
)

type reviewUserAPIPayload struct {
	Action string `json:"action"`
}

// @Summary review user
// @Tags user
// @Produce  json
// @Param body body updateUserAPIRequest true "update user API request body"
// @Success 200 {object} updateUserAPIResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /v1/management/user/review [patch]
// @Security basic
func review(
	router fiber.Router,
	authMiddleware authorization.AuthorizationMiddleware,
	userSqlAdapter userInterface.AdapterSQL,
	userService userInterface.Service,
) {
	routePath := fmt.Sprintf("%s/review/:id", prefix)
	requiredResources := []string{"USER_MANAGEMENT_REVIEW"}

	router.Patch(routePath, authMiddleware.Authorize(requiredResources), func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		modifier := string(c.Request().URI().Username())

		encodedUserId := c.Params("id", "")
		userId, err := publicfacingutil.Decode(encodedUserId)
		if err != nil {
			return err
		}

		var payload *reviewUserAPIPayload
		if err := c.BodyParser(&payload); err != nil {
			return errorcode.ErrCodeInvalidPayload
		}

		serviceDto, err := reviewUserPayloadValidation(ctx, userSqlAdapter, userId, payload, modifier)
		if err != nil {
			return err
		}

		err = userService.Review(ctx, serviceDto)
		if err != nil {
			return err
		}

		resp, err := compileGetDetailAPIResponse(ctx, userService, encodedUserId)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(resp)
	})
}

func reviewUserPayloadValidation(
	ctx context.Context,
	userSqlAdapter userInterface.AdapterSQL,
	userId int32,
	payload *reviewUserAPIPayload,
	modifier string,
) (*userDto.ServiceReviewPayload, error) {
	if payload.Action == "" {
		return nil, errorcode.ErrCodeInvalidStatus
	} else if !sliceutil.Contains([]reviewenum.Action{reviewenum.Approve, reviewenum.Reject}, reviewenum.Action(payload.Action)) {
		return nil, errorcode.ErrCodeInvalidStatus
	}

	existingUser, err := userSqlAdapter.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	notEligibleStatuses := []statusenum.Status{statusenum.Draft, statusenum.Rejected, statusenum.Active, statusenum.ActiveRejectSubmitted}
	if sliceutil.Contains(notEligibleStatuses, statusenum.Status(existingUser.Status)) {
		return nil, errorcode.WithCustomMessage(errorcode.ErrCodeInvalidStatus, "Invalid status")
	}

	return &userDto.ServiceReviewPayload{
		ExistingUser: existingUser,
		Action:       reviewenum.Action(payload.Action),
		Modifier:     modifier,
	}, nil

}
