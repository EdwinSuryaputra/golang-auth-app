package burequestbucket

import (
	"fmt"
	"strings"
	"time"

	"golang-auth-app/app/common/dto/request"
	"golang-auth-app/app/common/dto/response"

	timeutil "golang-auth-app/app/utils/time"

	burbInterface "golang-auth-app/app/interfaces/management/bu_request_bucket"
	burbDto "golang-auth-app/app/interfaces/management/bu_request_bucket/dto"

	"golang-auth-app/app/routes/rest/middleware/authorization"

	"github.com/gofiber/fiber/v2"
	"github.com/rotisserie/eris"
)

type getListCompletedAPIPayload struct {
	RequestDate string `json:"requestDate"`
	Username    string `json:"username"`
	Fullname    string `json:"fullname"`
	BuLevel     string `json:"buLevel"`
	BuLocation  string `json:"buLocation"`
	Status      string `json:"status"`
	ReviewedAt  string `json:"reviewedAt"`
	ReviewedBy  string `json:"reviewedBy"`
}

func getListCompleted(
	router fiber.Router,
	authMiddleware authorization.AuthorizationMiddleware,
	burbService burbInterface.Service,
) {
	routePath := fmt.Sprintf("%s/list/completed", prefix)
	requiredResources := []string{"NTE_BU_REQUEST_BUCKET_GET_LIST"}
	startTime := time.Now()

	router.Post(routePath, authMiddleware.Authorize(requiredResources), func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		payload, err := getListCompletedPayload(c)
		if err != nil {
			return err
		}

		data, err := burbService.GetListCompleted(ctx, payload)
		if err != nil {
			return err
		}

		var previousPage *int = nil
		if payload.Page-1 > 0 {
			previousPage = new(int)
			*previousPage = payload.Page - 1
		}

		return c.JSON(&response.GetListAPIResponse[*burbDto.ServiceGetListCompletedEntry]{
			DataList:    data.Entries,
			Pager:       response.SetPager(int64(data.TotalRow), payload.Page, payload.Limit),
			ProcessTime: response.GetProcessTime(startTime),
		})
	})
}

func getListCompletedPayload(c *fiber.Ctx) (*burbDto.ServiceGetListCompletedPayload, error) {
	var payload *request.GetListAPIRequest[getListCompletedAPIPayload]
	if err := c.BodyParser(&payload); err != nil {
		return nil, eris.Wrap(err, err.Error())
	}

	requestDate, err := timeutil.ParseDateFromString(strings.TrimSpace(payload.PropertyFilter.RequestDate))
	if err != nil {
		return nil, err
	}

	reviewedAt, err := timeutil.ParseDateFromString(strings.TrimSpace(payload.PropertyFilter.ReviewedAt))
	if err != nil {
		return nil, err
	}

	page := 1
	if payload.Page != 0 {
		page = payload.Page
	}

	limit := 10
	if payload.Limit != 0 {
		limit = payload.Limit
	}

	return &burbDto.ServiceGetListCompletedPayload{
		RequestDate: requestDate,
		Username:    strings.TrimSpace(payload.PropertyFilter.Username),
		Fullname:    strings.TrimSpace(payload.PropertyFilter.Fullname),
		BuLevel:     strings.TrimSpace(payload.PropertyFilter.BuLevel),
		BuLocation:  strings.TrimSpace(payload.PropertyFilter.BuLocation),
		Status:      strings.TrimSpace(payload.PropertyFilter.Status),
		ReviewedBy:  strings.TrimSpace(payload.PropertyFilter.ReviewedBy),
		ReviewedAt:  reviewedAt,
		Page:        page,
		Limit:       limit,
	}, nil
}
