package user

import (
	"fmt"
	"strings"
	"time"

	"golang-auth-app/app/common/dto/request"
	"golang-auth-app/app/common/dto/response"

	"golang-auth-app/app/interfaces/errorcode"

	userInterface "golang-auth-app/app/interfaces/management/user"
	userDto "golang-auth-app/app/interfaces/management/user/dto"

	"golang-auth-app/app/routes/rest/middleware/authorization"

	publicfacingutil "golang-auth-app/app/utils/publicfacing"
	sliceutil "golang-auth-app/app/utils/slice"
	timeutil "golang-auth-app/app/utils/time"

	"github.com/gofiber/fiber/v2"
	"github.com/rotisserie/eris"
)

type getListAPIRequestPropertyFilter struct {
	Username     string `json:"username"`
	Type         string `json:"type"`
	SupplierName string `json:"supplierName"`
	BuLevel      string `json:"buLevel"`
	BuLocation   string `json:"buLocation"`
	Status       string `json:"status"`
	StartDate    string `json:"startDate"`
	EndDate      string `json:"endDate"`
	CreatedBy    string `json:"createdBy"`
	CreatedAt    string `json:"createdAt"`
	UpdatedBy    string `json:"updatedBy"`
	UpdatedAt    string `json:"updatedAt"`
}

type getListAPIResponseEntry struct {
	Id           string     `json:"id"`
	Username     string     `json:"username"`
	Type         string     `json:"type"`
	SupplierName *string    `json:"supplierName"`
	BuLevel      *string    `json:"buLevel"`
	BuLocation   *string    `json:"buLocation"`
	Status       string     `json:"status"`
	StartDate    *time.Time `json:"startDate"`
	EndDate      *time.Time `json:"endDate"`
	CreatedAt    time.Time  `json:"createdAt"`
	CreatedBy    string     `json:"createdBy"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	UpdatedBy    string     `json:"updatedBy"`
}

// @Summary Get list users
// @Tags user
// @Produce  json
// @Param body body updateUserAPIRequest true "update user API request body"
// @Success 200 {object} updateUserAPIResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /v1/management/user/list [get]
// @Security basic
func getList(
	router fiber.Router,
	authMiddleware authorization.AuthorizationMiddleware,
	userService userInterface.Service,
) {
	routePath := fmt.Sprintf("%s/list", prefix)
	requiredResources := []string{"NTE_USER_MANAGEMENT_GET_LIST"}
	startTime := time.Now()

	router.Post(routePath, authMiddleware.Authorize(requiredResources), func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		payload, err := getListPayload(c)
		if err != nil {
			return err
		}

		usersPaginated, err := userService.GetList(ctx, payload)
		if err != nil {
			return err
		}

		var previousPage *int = nil
		if payload.Query.Page-1 > 0 {
			previousPage = new(int)
			*previousPage = payload.Query.Page - 1
		}

		entries, err := sliceutil.MapWithError(usersPaginated.Entries, func(dt *userDto.AdapterSqlGetPaginatedEntry) (getListAPIResponseEntry, error) {
			encodedId, err := publicfacingutil.Encode(dt.ID)
			if err != nil {
				return getListAPIResponseEntry{}, err
			}

			return getListAPIResponseEntry{
				Id:           encodedId,
				Username:     dt.Username,
				Type:         dt.Type,
				SupplierName: dt.SupplierName,
				BuLevel:      dt.BuLevel,
				BuLocation:   dt.BuLocation,
				Status:       dt.Status,
				StartDate:    dt.StartDate,
				EndDate:      dt.EndDate,
				CreatedAt:    dt.CreatedAt,
				CreatedBy:    dt.CreatedBy,
				UpdatedAt:    dt.UpdatedAt,
				UpdatedBy:    dt.UpdatedBy,
			}, nil
		})
		if err != nil {
			return err
		}

		return c.JSON(response.GetListAPIResponse[getListAPIResponseEntry]{
			DataList:    entries,
			Pager:       response.SetPager(int64(usersPaginated.TotalRow), payload.Query.Page, payload.Query.Limit),
			ProcessTime: response.GetProcessTime(startTime),
		})
	})
}

func getListPayload(c *fiber.Ctx) (*userDto.ServiceGetListPayload, error) {
	var payload *request.GetListAPIRequest[getListAPIRequestPropertyFilter]
	if err := c.BodyParser(&payload); err != nil {
		return nil, eris.Wrap(err, err.Error())
	}

	startDate, err := timeutil.ParseDateFromString(strings.TrimSpace(payload.PropertyFilter.StartDate))
	if err != nil {
		return nil, err
	}

	endDate, err := timeutil.ParseDateFromString(strings.TrimSpace(payload.PropertyFilter.EndDate))
	if err != nil {
		return nil, err
	}

	if startDate != nil && endDate != nil && endDate.Before(*startDate) {
		return nil, errorcode.ErrCodeEndDateLesserThanStartDate
	}

	createdAt, err := timeutil.ParseDateFromString(strings.TrimSpace(payload.PropertyFilter.CreatedAt))
	if err != nil {
		return nil, err
	}

	updatedAt, err := timeutil.ParseDateFromString(strings.TrimSpace(payload.PropertyFilter.UpdatedAt))
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

	return &userDto.ServiceGetListPayload{
		Query: &userDto.AdapterSqlGetPaginatedPayload{
			Username:     strings.TrimSpace(payload.PropertyFilter.Username),
			Type:         strings.TrimSpace(payload.PropertyFilter.Type),
			SupplierName: strings.TrimSpace(payload.PropertyFilter.SupplierName),
			BuLevel:      strings.TrimSpace(payload.PropertyFilter.BuLevel),
			BuLocation:   strings.TrimSpace(payload.PropertyFilter.BuLocation),
			StartDate:    startDate,
			EndDate:      endDate,
			Status:       strings.TrimSpace(payload.PropertyFilter.Status),
			CreatedBy:    strings.TrimSpace(payload.PropertyFilter.CreatedBy),
			CreatedAt:    createdAt,
			UpdatedBy:    strings.TrimSpace(payload.PropertyFilter.UpdatedBy),
			UpdatedAt:    updatedAt,
			Page:         page,
			Limit:        limit,
		},
	}, nil
}
