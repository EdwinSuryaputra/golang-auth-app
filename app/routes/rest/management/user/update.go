package user

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-auth-app/config"
	"time"

	statusenum "golang-auth-app/app/common/enums/status"

	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/common/errorcode"

	roleInterface "golang-auth-app/app/interfaces/management/role"
	userInterface "golang-auth-app/app/interfaces/management/user"
	userDto "golang-auth-app/app/interfaces/management/user/dto"

	"golang-auth-app/app/routes/rest/middleware/authorization"

	publicfacingutil "golang-auth-app/app/utils/publicfacing"
	sliceutil "golang-auth-app/app/utils/slice"

	"github.com/gofiber/fiber/v2"
	"github.com/rotisserie/eris"
)

type updateUserAPIPayload struct {
	IsInactive      bool     `json:"isInactive"`
	IsSubmit        bool     `json:"isSubmit"`
	Username        string   `json:"username"`
	FullName        string   `json:"fullName"`
	Email           string   `json:"email"`
	AssignedRoleIds []string `json:"assignedRoleIds"`
}

// @Summary update a user
// @Tags user
// @Produce  json
// @Param body body updateUserAPIRequest true "update user API request body"
// @Success 200 {object} updateUserAPIResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /v1/management/user [put]
// @Security basic
func update(
	router fiber.Router,
	authMiddleware authorization.AuthorizationMiddleware,
	roleSqlAdapter roleInterface.AdapterSQL,
	userSqlAdapter userInterface.AdapterSQL,
	userService userInterface.Service,
) {
	routePath := fmt.Sprintf("%s/:id", prefix)
	requiredResources := []string{"USER_MANAGEMENT_UPDATE"}

	router.Put(routePath, authMiddleware.Authorize(requiredResources), func(c *fiber.Ctx) error {
		ctx := c.UserContext()
		modifier := string(c.Request().URI().Username())

		encodedUserId := c.Params("id", "")
		userId, err := publicfacingutil.Decode(encodedUserId)
		if err != nil {
			return err
		}

		var payload *updateUserAPIPayload
		if err := c.BodyParser(&payload); err != nil {
			return errorcode.ErrCodeInvalidPayload
		}

		serviceDto, err := updateUserPayloadValidation(ctx, roleSqlAdapter, userSqlAdapter, userId, payload, modifier)
		if err != nil {
			return err
		}

		err = userService.Update(ctx, serviceDto)
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

func updateUserPayloadValidation(
	ctx context.Context,
	roleSqlAdapter roleInterface.AdapterSQL,
	userSqlAdapter userInterface.AdapterSQL,
	userId int32,
	payload *updateUserAPIPayload,
	modifier string,
) (*userDto.ServiceUpdatePayload, error) {
	now := time.Now()

	currentDataUser, err := userSqlAdapter.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	newDataUser := *currentDataUser
	newDataUser.UpdatedAt = now
	newDataUser.UpdatedBy = modifier

	eligibleStatuses := []statusenum.Status{statusenum.Draft, statusenum.Rejected, statusenum.Active, statusenum.ActiveRejectSubmitted}
	if !sliceutil.Contains(eligibleStatuses, statusenum.Status(currentDataUser.Status)) {
		return nil, errorcode.WithCustomMessage(errorcode.ErrCodeInvalidStatus, "Invalid status")
	}

	if payload.IsInactive {
		return updateUserPayloadInactiveValidation(currentDataUser, &newDataUser)
	}

	if err = updateUserPayloadBasicDataValidation(payload, currentDataUser, &newDataUser); err != nil {
		return nil, err
	}

	if err = updateUserPayloadNewStatusAssignment(payload, currentDataUser, &newDataUser); err != nil {
		return nil, err
	}

	if err = updateUserPayloadRoleValidation(ctx, roleSqlAdapter, payload, currentDataUser, &newDataUser, now, modifier); err != nil {
		return nil, err
	}

	return &userDto.ServiceUpdatePayload{
		CurrentDataUser: currentDataUser,
		NewDataUser:     &newDataUser,
	}, nil
}

func updateUserPayloadInactiveValidation(
	currentDataUser *model.User,
	newDataUser *model.User,
) (*userDto.ServiceUpdatePayload, error) {
	if currentDataUser.Status != string(statusenum.Active) {
		return nil, errorcode.WithCustomMessage(errorcode.ErrCodeInvalidStatus, "Unable to proceed due to document status is not active yet")
	}

	newDataUser.Status = string(statusenum.ActiveInactivationSubmitted)

	return &userDto.ServiceUpdatePayload{
		CurrentDataUser: currentDataUser,
		NewDataUser:     newDataUser,
	}, nil
}

func updateUserPayloadBasicDataValidation(
	payload *updateUserAPIPayload,
	currentDataUser *model.User,
	newDataUser *model.User,
) error {
	cfg := config.Module.User.PayloadValidation

	if !sliceutil.Contains([]statusenum.Status{statusenum.Draft, statusenum.Rejected}, statusenum.Status(currentDataUser.Status)) {
		if payload.Username == "" {
			return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest, "Username is required")
		} else if len(payload.Username) < cfg.UsernameMinDigit {
			return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest,
				fmt.Sprintf("Username's min length is %d digit", cfg.UsernameMinDigit))
		}

		if payload.FullName == "" {
			return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest, "Full name is required")
		} else if len(payload.FullName) < cfg.FullNameMinDigit {
			return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest, fmt.Sprintf("Full name's min length is %d digit", cfg.FullNameMinDigit))
		}

		if payload.Email == "" {
			return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest, "Email is required")
		} else if len(payload.Email) < cfg.EmailMinDigit {
			return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest, fmt.Sprintf("Email's min length is %d digit", cfg.EmailMinDigit))
		}

		newDataUser.Username = payload.Username
		newDataUser.FullName = payload.FullName
		newDataUser.Email = payload.Email
	}

	return nil
}

func updateUserPayloadRoleValidation(
	ctx context.Context,
	roleSqlAdapter roleInterface.AdapterSQL,
	payload *updateUserAPIPayload,
	currentDataUser *model.User,
	newDataUser *model.User,
	now time.Time,
	modifier string,
) error {
	if len(payload.AssignedRoleIds) < 1 {
		return errorcode.WithCustomMessage(errorcode.ErrCodeBadRequest, "Assigned role ids is required")
	} else {
		inputRoleIds, err := sliceutil.MapWithError(payload.AssignedRoleIds, func(roleId string) (int32, error) {
			return publicfacingutil.Decode(roleId)
		})
		if err != nil {
			return err
		}

		existingRoles, err := roleSqlAdapter.GetRolesByIds(ctx, inputRoleIds)
		if err != nil {
			return err
		} else if len(existingRoles) < 1 {
			return errorcode.WithCustomMessage(errorcode.ErrCodeRoleNotFound, "Roles are not found")
		}
		existingRolesByIds := sliceutil.AssociateBy(existingRoles, func(dt *model.Role) int32 { return dt.ID })

		newUserRoles := []*model.UserRoleMapping{}
		for _, roleId := range inputRoleIds {
			existingRole, isExist := existingRolesByIds[roleId]
			if !isExist {
				return errorcode.WithCustomMessage(errorcode.ErrCodeRoleNotFound, fmt.Sprintf("Role id %d is not found", roleId))
			} else if existingRole.Type != currentDataUser.Type {
				return errorcode.ErrCodeInvalidRoleType
			}

			newUserRoles = append(newUserRoles, &model.UserRoleMapping{
				UserID:    &currentDataUser.ID,
				RoleID:    &existingRole.ID,
				CreatedBy: modifier,
				CreatedAt: now,
				UpdatedBy: modifier,
				UpdatedAt: now,
			})
		}

		marshalledNewUserRole, err := json.Marshal(newUserRoles)
		if err != nil {
			return eris.Wrap(err, "error occurred during stringify resources")
		}
		assignedRoles := string(marshalledNewUserRole)

		newDataUser.AssignedRoles = &assignedRoles
	}

	return nil
}

func updateUserPayloadNewStatusAssignment(
	payload *updateUserAPIPayload,
	currentDataUser *model.User,
	newDataUser *model.User,
) error {
	var newStatus statusenum.Status
	switch payload.IsSubmit {
	case true:
		switch statusenum.Status(currentDataUser.Status) {
		case statusenum.Draft:
			newStatus = statusenum.Submitted
		case statusenum.Active:
			newStatus = statusenum.ActiveEditSubmitted
		}
	case false:
		if currentDataUser.Status != string(statusenum.Draft) {
			return errorcode.ErrCodeInvalidStatus
		}
		newStatus = statusenum.Draft
	}
	newDataUser.Status = newStatus.ToString()

	return nil
}
