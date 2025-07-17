package role

import (
	"fmt"
	"strings"
	"time"

	"golang-auth-app/app/common/dto/request"
	"golang-auth-app/app/common/dto/response"

	"golang-auth-app/app/datasources/sql/gorm/model"

	roleInterface "golang-auth-app/app/interfaces/management/role"

	"golang-auth-app/app/routes/rest/middleware/authorization"

	publicfacingutil "golang-auth-app/app/utils/publicfacing"
	sliceutil "golang-auth-app/app/utils/slice"
	timeutil "golang-auth-app/app/utils/time"

	"github.com/gofiber/fiber/v2"
	"github.com/rotisserie/eris"
)

type getListAPIRequestPropertyFilter struct {
	Name         string `json:"name"`
	Status       string `json:"status"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	InactiveDate string `json:"inactiveDate"`
	CreatedBy    string `json:"createdBy"`
	CreatedAt    string `json:"createdAt"`
	UpdatedBy    string `json:"updatedBy"`
	UpdatedAt    string `json:"updatedAt"`
}

type getListAPIResponseEntry struct {
	Id           string     `json:"id"`
	Name         string     `json:"name"`
	Status       string     `json:"status"`
	Description  string     `json:"description"`
	Type         string     `json:"type"`
	InactiveDate *time.Time `json:"inactiveDate"`
	CreatedAt    time.Time  `json:"createdAt"`
	CreatedBy    string     `json:"createdBy"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	UpdatedBy    string     `json:"updatedBy"`
}

func getList(
	router fiber.Router,
	authMiddleware authorization.AuthorizationMiddleware,
	roleService roleInterface.Service,
) {
	routePath := fmt.Sprintf("%s/list", prefix)
	requiredResources := []string{"NTE_ROLE_MANAGEMENT_GET_LIST"}
	startTime := time.Now()

	router.Post(routePath, authMiddleware.Authorize(requiredResources), func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		payload, err := getListPayload(c)
		if err != nil {
			return err
		}

		roles, err := roleService.GetList(ctx, payload)
		if err != nil {
			return err
		}

		var previousPage *int = nil
		if payload.Page-1 > 0 {
			previousPage = new(int)
			*previousPage = payload.Page - 1
		}

		entries, err := sliceutil.MapWithError(roles.Entries, func(dt *model.Role) (getListAPIResponseEntry, error) {
			encodedId, err := publicfacingutil.Encode(dt.ID)
			if err != nil {
				return getListAPIResponseEntry{}, err
			}

			return getListAPIResponseEntry{
				Id:           encodedId,
				Name:         dt.Name,
				Status:       dt.Status,
				Description:  dt.Description,
				Type:         dt.Type,
				InactiveDate: dt.InactiveDate,
				CreatedAt:    dt.CreatedAt,
				CreatedBy:    dt.CreatedBy,
				UpdatedAt:    dt.UpdatedAt,
				UpdatedBy:    dt.UpdatedBy,
			}, nil
		})
		if err != nil {
			return err
		}

		return c.JSON(&response.GetListAPIResponse[getListAPIResponseEntry]{
			DataList:    entries,
			Pager:       response.SetPager(int64(roles.TotalRow), payload.Page, payload.Limit),
			ProcessTime: response.GetProcessTime(startTime),
		})
	})
}

func getListPayload(c *fiber.Ctx) (*roleInterface.AdapterSqlGetPaginatedPayload, error) {
	var payload *request.GetListAPIRequest[getListAPIRequestPropertyFilter]
	if err := c.BodyParser(&payload); err != nil {
		return nil, eris.Wrap(err, err.Error())
	}

	inactiveDate, err := timeutil.ParseDateFromString(strings.TrimSpace(payload.PropertyFilter.InactiveDate))
	if err != nil {
		return nil, err
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

	return &roleInterface.AdapterSqlGetPaginatedPayload{
		Query: &roleInterface.GetPaginatedColumnsFilter{
			Name:         strings.TrimSpace(payload.PropertyFilter.Name),
			Status:       strings.TrimSpace(payload.PropertyFilter.Status),
			Description:  strings.TrimSpace(payload.PropertyFilter.Description),
			Type:         strings.TrimSpace(payload.PropertyFilter.Type),
			InactiveDate: inactiveDate,
			CreatedBy:    strings.TrimSpace(payload.PropertyFilter.CreatedBy),
			CreatedAt:    createdAt,
			UpdatedBy:    strings.TrimSpace(payload.PropertyFilter.UpdatedBy),
			UpdatedAt:    updatedAt,
		},
		Page:  page,
		Limit: limit,
	}, nil
}
