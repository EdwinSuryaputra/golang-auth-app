package errorcode

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ErrCodeInternalServerError        errorItf = Error{fiber.StatusInternalServerError, "INTERNAL_SERVER_ERROR", "Something went wrong"}
	ErrCodeInvalidInput               errorItf = Error{fiber.StatusBadRequest, "INVALID_INPUT", "Invalid input"}
	ErrCodeNotFound                   errorItf = Error{fiber.StatusNotFound, "NOT_FOUND", "Data not found"}
	ErrCodeUnauthorized               errorItf = Error{fiber.StatusUnauthorized, "UNAUTHORIZED", "Unauthorized"}
	ErrCodeForbidden                  errorItf = Error{fiber.StatusForbidden, "FORBIDDEN", "Forbidden request"}
	ErrCodeBadRequest                 errorItf = Error{fiber.StatusBadRequest, "BAD_REQUEST", "Bad request"}
	ErrCodeUnprocessableEntity        errorItf = Error{fiber.StatusUnprocessableEntity, "UNPROCESSABLE_ENTITY", "Unprocessable entity"}
	ErrCodeInvalidPayload             errorItf = Error{fiber.StatusBadRequest, "INVALID_PAYLOAD", "Invalid payload"}
	ErrCodeInvalidParameter           errorItf = Error{fiber.StatusBadRequest, "INVALID_PARAMETER", "Invalid parameter"}
	ErrCodeInvalidId                  errorItf = Error{fiber.StatusBadRequest, "INVALID_ID", "Invalid id"}
	ErrCodeInvalidHeader              errorItf = Error{fiber.StatusBadRequest, "INVALID_HEADER", "Invalid header"}
	ErrCodeMissingAction              errorItf = Error{fiber.StatusBadRequest, "MISSING_ACTION", "Missing action"}
	ErrCodeInvalidAction              errorItf = Error{fiber.StatusBadRequest, "INVALID_ACTION", "Invalid action"}
	ErrCodeInvalidStatus              errorItf = Error{fiber.StatusBadRequest, "INVALID_STATUS", "Invalid status"}
	ErrCodeStatusRequired             errorItf = Error{fiber.StatusBadRequest, "STATUS_REQUIRED", "Status is required"}
	ErrCodeStatusIsNotSubmitted       errorItf = Error{fiber.StatusUnprocessableEntity, "STATUS_NOT_SUBMITTED", "Status is not submitted yet"}
	ErrCodeStatusIsAlreadyApproved    errorItf = Error{fiber.StatusUnprocessableEntity, "STATUS_ALREADY_APPROVED", "Status is already approved"}
	ErrCodeStatusIsAlreadyRejected    errorItf = Error{fiber.StatusUnprocessableEntity, "STATUS_ALREADY_REJECTED", "Status is already rejected"}
	ErrCodeDetailsRequired            errorItf = Error{fiber.StatusBadRequest, "DETAILS_REQUIRED", "Details is required"}
	ErrCodeEndDateLesserThanStartDate errorItf = Error{fiber.StatusBadRequest, "END_DATE_LESSER_THAN_START_DATE", "End date can't be lesser than start date"}
)

// Authentication & Authorization
var (
	ErrCodeInvalidUsername     errorItf = Error{fiber.StatusUnauthorized, "INVALID_USERNAME", "Invalid username"}
	ErrCodeInvalidPassword     errorItf = Error{fiber.StatusUnauthorized, "INVALID_PASSWORD", "Invalid password"}
	ErrCodeInvalidOldPassword  errorItf = Error{fiber.StatusUnauthorized, "INVALID_OLD_PASSWORD", "Invalid old password"}
	ErrCodeMissingAuthToken    errorItf = Error{fiber.StatusUnauthorized, "MISSING_AUTH_TOKEN", "Missing auth token"}
	ErrCodeMissingApplication  errorItf = Error{fiber.StatusUnauthorized, "MISSING_APPLICATION", "Missing application"}
	ErrCodeUserHasNoRole       errorItf = Error{fiber.StatusForbidden, "USER_HAS_NO_ROLE", "User has no role"}
	ErrCodeInvalidTokenFormat  errorItf = Error{fiber.StatusUnauthorized, "INVALID_AUTH_TOKEN_FORMAT", "Invalid authentication token format"}
	ErrCodeInvalidTokenExpired errorItf = Error{fiber.StatusUnauthorized, "AUTH_TOKEN_EXPIRED", "Authentication token expired"}
	ErrCodeNoPermissionAccess  errorItf = Error{fiber.StatusForbidden, "NO_PERMISSION_ACCESS", "You don't have permission to access this resource"}
)

var (
	ErrCodeApplicationNotFound errorItf = Error{fiber.StatusNotFound, "APPLICATION_NOT_FOUND", "Application not found"}
)

// Role
var (
	ErrCodeResourcesRequired errorItf = Error{fiber.StatusBadRequest, "RESOURCES_REQUIRED", "Menu assignment is required"}
	ErrCodeFunctionsRequired errorItf = Error{fiber.StatusBadRequest, "FUNCTIONS_REQUIRED", "Function assignment is required"}
)

var (
	ErrCodeRoleNotFound    errorItf = Error{fiber.StatusNotFound, "ROLE_NOT_FOUND", "Role not found"}
	ErrCodeInvalidRoleType errorItf = Error{fiber.StatusBadRequest, "INVALID_ROLE_TYPE", "Invalid role type"}
	ErrCodeInvalidMenu     errorItf = Error{fiber.StatusBadRequest, "INVALID_MENU", "Invalid menu"}
	ErrCodeInvalidSubmenu  errorItf = Error{fiber.StatusBadRequest, "INVALID_SUBMENU", "Invalid sub menu"}
	ErrCodeInvalidFunction errorItf = Error{fiber.StatusBadRequest, "INVALID_FUNCTION", "Invalid function"}
)

var (
	ErrCodeUserNotFound     errorItf = Error{fiber.StatusNotFound, "USER_NOT_FOUND", "User not found"}
	ErrCodeUserAlreadyExist errorItf = Error{fiber.StatusUnprocessableEntity, "USER_ALREADY_EXIST", "User already exist"}
	ErrCodeUserNotActive    errorItf = Error{fiber.StatusUnauthorized, "USER_NOT_ACTIVE", "User not active"}
	ErrCodeMissingUserType  errorItf = Error{fiber.StatusBadRequest, "MISSING_USER_TYPE", "Missing user type"}
	ErrCodeInvalidUserType  errorItf = Error{fiber.StatusBadRequest, "INVALID_USER_TYPE", "Invalid user type"}
)

var (
	ErrCodeBurbNotFound errorItf = Error{fiber.StatusNotFound, "BURB_NOT_FOUND", "BU request bucket data is not found"}
)

var (
	ErrCodeCasbinRuleAlreadyExist errorItf = Error{fiber.StatusUnprocessableEntity, "CASBIN_RULE_ALREADY_EXIST", "Casbin rule already exist"}
)

var (
	ErrCodeMissingRequiredResource errorItf = Error{fiber.StatusBadRequest, "MISSING_REQUIRED_RESOURCE", "Required resource is required"}
)
