package response

import (
	"encoding/json"
	"fmt"
	"runtime/debug"

	loggerenum "golang-auth-app/app/common/enums/logger"
	"golang-auth-app/app/interfaces/errorcode"

	"github.com/gofiber/fiber/v2"
	"github.com/rotisserie/eris"
	"go.uber.org/zap"
)

type ResponseSuccessWrapper[T any] struct {
	Data T `json:"data"`
}

type ResponseErrorWrapper struct {
	Error ErrorResponse `json:"error"`
}

type ErrorResponse struct {
	ErrEnum string `json:"enum"`
	Message string `json:"message"`
}

type Wrapper struct {
	Handler fiber.Handler
}

func New(logger *zap.Logger) Wrapper {
	handler := func(c *fiber.Ctx) error {
		var requestId string
		if reqId, isExist := c.Locals("requestId").(string); isExist {
			requestId = reqId
		}

		defer func() {
			if r := recover(); r != nil {
				handlePanic(c, r, logger, requestId)
			}
		}()

		err := c.Next()
		if err != nil {
			return handleError(c, err, logger, requestId)
		}

		return handleSuccess(c, logger, requestId)
	}

	return Wrapper{Handler: handler}
}

func handleSuccess(c *fiber.Ctx, logger *zap.Logger, requestId string) error {
	zapFields := []zap.Field{
		zap.String(loggerenum.RequestId.ToString(), requestId),
		zap.String(loggerenum.Path.ToString(), c.Path()),
		zap.String(loggerenum.Method.ToString(), c.Method()),
		zap.String(loggerenum.Ip.ToString(), c.IP()),
		zap.Int(loggerenum.StatusCode.ToString(), c.Response().StatusCode()),
	}

	responseBody := c.Response().Body()
	if len(responseBody) > 0 {
		zapFields = append(zapFields, zap.Any(loggerenum.ResponseBody.ToString(), responseBody))
	}

	logger.Info(loggerenum.APIResponse.ToString(), zapFields...)

	return c.JSON(json.RawMessage(responseBody))
}

func handleError(c *fiber.Ctx, err error, logger *zap.Logger, requestId string) error {
	zapFields := []zap.Field{
		zap.String(loggerenum.RequestId.ToString(), requestId),
		zap.String(loggerenum.Path.ToString(), c.Path()),
		zap.String(loggerenum.Method.ToString(), c.Method()),
		zap.String(loggerenum.Ip.ToString(), c.IP()),
	}

	responseBody := c.Response().Body()
	if len(responseBody) > 0 {
		zapFields = append(zapFields, zap.Any(loggerenum.ResponseBody.ToString(), responseBody))
	}

	errRoot := eris.Unpack(err).ErrRoot
	if len(errRoot.Stack) > 0 {
		frame := errRoot.Stack[0]
		zapFields = append(zapFields, zap.String(loggerenum.Stacktrace.ToString(), fmt.Sprintf("%s:%d %s", frame.File, frame.Line, errRoot.Msg)))
	}

	var errObj errorcode.Error
	if eris.As(err, &errObj) {
		errResp := &ResponseErrorWrapper{
			Error: ErrorResponse{
				ErrEnum: errObj.GetErrEnum(),
				Message: errObj.Error(),
			},
		}

		zapFields = append(zapFields,
			zap.Int(loggerenum.StatusCode.ToString(), errObj.GetHttpStatusCode()),
			zap.Any(loggerenum.ResponseBody.ToString(), errResp),
		)

		if errObj.ErrHttpStatusCode >= 500 {
			logger.Error(loggerenum.APIResponse.ToString(), zapFields...)
		} else if errObj.ErrHttpStatusCode >= 400 {
			logger.Warn(loggerenum.APIResponse.ToString(), zapFields...)
		}

		return c.Status(errObj.GetHttpStatusCode()).JSON(errResp)
	}

	zapFields = append(zapFields, zap.Int(loggerenum.StatusCode.ToString(), errorcode.ErrCodeInternalServerError.GetHttpStatusCode()))

	logger.Error(loggerenum.APIResponse.ToString(), zapFields...)

	return c.Status(errorcode.ErrCodeInternalServerError.GetHttpStatusCode()).JSON(ResponseErrorWrapper{
		Error: ErrorResponse{
			ErrEnum: errorcode.ErrCodeInternalServerError.GetErrEnum(),
			Message: err.Error(),
		},
	})
}

func handlePanic(c *fiber.Ctx, r any, logger *zap.Logger, requestId string) error {
	var err error
	switch x := r.(type) {
	case string:
		err = fmt.Errorf("%s", x)
	case error:
		err = x
	default:
		err = fmt.Errorf("%v", x)
	}

	zapFields := []zap.Field{
		zap.String(loggerenum.RequestId.ToString(), requestId),
		zap.String(loggerenum.Path.ToString(), c.Path()),
		zap.String(loggerenum.Method.ToString(), c.Method()),
		zap.String(loggerenum.Ip.ToString(), c.IP()),
		zap.String(loggerenum.Panic.ToString(), err.Error()),
		zap.ByteString(loggerenum.Stacktrace.ToString(), debug.Stack()),
	}

	logger.Error(loggerenum.APIResponse.ToString(), zapFields...)

	return c.Status(errorcode.ErrCodeInternalServerError.GetHttpStatusCode()).JSON(ResponseErrorWrapper{
		Error: ErrorResponse{
			ErrEnum: errorcode.ErrCodeInternalServerError.GetErrEnum(),
			Message: errorcode.ErrCodeInternalServerError.Error(),
		},
	})
}
