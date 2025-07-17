package burequestbucket

import (
	"fmt"

	reviewenum "golang-auth-app/app/common/enums/review"

	publicfacingutil "golang-auth-app/app/utils/publicfacing"
	sliceutil "golang-auth-app/app/utils/slice"

	"golang-auth-app/app/interfaces/errorcode"

	burbInterface "golang-auth-app/app/interfaces/management/bu_request_bucket"
	burbDto "golang-auth-app/app/interfaces/management/bu_request_bucket/dto"

	"golang-auth-app/app/routes/rest/middleware/authorization"

	"github.com/gofiber/fiber/v2"
	"github.com/rotisserie/eris"
)

type reviewAPIPayload struct {
	Action string `json:"action"`
}

func review(
	router fiber.Router,
	authMiddleware authorization.AuthorizationMiddleware,
	burbService burbInterface.Service,
) {
	routePath := fmt.Sprintf("%s/review/:burbId", prefix)
	requiredResources := []string{"NTE_BU_REQUEST_BUCKET_REVIEW"}

	router.Patch(routePath, authMiddleware.Authorize(requiredResources), func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		modifier := string(c.Request().URI().Username())

		encodedBuId := c.Params("burbId", "")
		burbId, err := publicfacingutil.Decode(encodedBuId)
		if err != nil {
			return eris.Wrap(err, "error occurred during decode public id")
		}

		payload, err := reviewPayload(c, burbId, modifier)
		if err != nil {
			return err
		}

		err = burbService.Review(ctx, payload)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(map[string]any{
			"result": true,
		})
	})
}

func reviewPayload(c *fiber.Ctx, burbId int32, modifier string) (*burbDto.ServiceReviewPayload, error) {
	var payload *reviewAPIPayload
	if err := c.BodyParser(&payload); err != nil {
		return nil, eris.Wrap(err, err.Error())
	}

	if payload.Action == "" {
		return nil, errorcode.ErrCodeMissingAction
	} else if !sliceutil.Contains([]reviewenum.Action{reviewenum.Approve, reviewenum.Reject}, reviewenum.Action(payload.Action)) {
		return nil, errorcode.ErrCodeInvalidAction
	}

	return &burbDto.ServiceReviewPayload{
		Id:       burbId,
		Action:   reviewenum.Action(payload.Action),
		Modifier: modifier,
	}, nil
}
