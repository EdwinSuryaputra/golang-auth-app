package user

import (
	"fmt"
	"golang-auth-app/config"
	"time"

	statusenum "golang-auth-app/app/common/enums/status"
	userenum "golang-auth-app/app/common/enums/user"

	"golang-auth-app/app/datasources/sql/gorm/model"
	"golang-auth-app/app/interfaces/errorcode"
	"golang-auth-app/app/interfaces/management/user"
	"golang-auth-app/app/routes/rest/middleware/authorization"

	publicfacingutil "golang-auth-app/app/utils/publicfacing"

	"github.com/gofiber/fiber/v2"
)

type createDraftUserAPIPayload struct {
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Type     string `json:"type"`
}

// @Summary Create a user
// @Tags user
// @Produce  json
// @Param body body createUserAPIRequest true "Create user API request body"
// @Success 200 {object} createUserAPIResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /v1/management/user [post]
// @Security basic
func createDraft(
	router fiber.Router,
	authMiddleware authorization.AuthorizationMiddleware,
	userSqlAdapter user.AdapterSQL,
	userService user.Service,
) {
	routePath := fmt.Sprintf("%s/", prefix)
	requiredResources := []string{"NTE_USER_MANAGEMENT_CREATE_DRAFT"}

	router.Post(routePath, authMiddleware.Authorize(requiredResources), func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		var payload *createDraftUserAPIPayload
		if err := c.BodyParser(&payload); err != nil {
			return errorcode.ErrCodeInvalidPayload
		} else if err := createUserPayloadValidation(payload); err != nil {
			return err
		}

		modifier := string(c.Request().URI().Username())
		createdAt := time.Now()

		if existingUser, err := userSqlAdapter.GetUserByUsername(ctx, payload.Username); err != nil {
			return err
		} else if existingUser != nil {
			return errorcode.ErrCodeUserAlreadyExist
		}

		user := &model.User{
			Username:    payload.Username,
			Description: "",
			FullName:    payload.FullName,
			Email:       payload.Email,
			Password:    "",
			Type:        payload.Type,
			Status:      string(statusenum.Draft),
			CreatedBy:   modifier,
			CreatedAt:   createdAt,
			UpdatedBy:   modifier,
			UpdatedAt:   createdAt,
		}
		err := userService.CreateDraft(ctx, user)
		if err != nil {
			return err
		}

		encodedUserId, err := publicfacingutil.Encode(user.ID)
		if err != nil {
			return err
		}

		resp, err := compileGetDetailAPIResponse(ctx, userService, encodedUserId)
		if err != nil {
			return err
		}

		return c.Status(fiber.StatusCreated).JSON(resp)
	})
}

func createUserPayloadValidation(payload *createDraftUserAPIPayload) error {
	cfg := config.Module.User.PayloadValidation

	if payload.Username == "" {
		return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest, "Username is required")
	} else if len(payload.Username) < cfg.UsernameMinDigit {
		return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest,
			fmt.Sprintf("Username's min length is %d digit", cfg.UsernameMinDigit))
	}

	if payload.FullName == "" {
		return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest, "Full name is required")
	} else if len(payload.FullName) < cfg.FullNameMinDigit {
		return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest,
			fmt.Sprintf("Full name's min length is %d digit", cfg.FullNameMinDigit))
	}

	if payload.Email == "" {
		return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest, "Email is required")
	} else if len(payload.Email) < cfg.EmailMinDigit {
		return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest,
			fmt.Sprintf("Email's min length is %d digit", cfg.EmailMinDigit))
	}

	if payload.Type == "" {
		return errorcode.ErrCodeMissingUserType
	} else if !userenum.IsValidType(payload.Type) {
		return errorcode.ErrCodeInvalidUserType
	}

	return nil
}
