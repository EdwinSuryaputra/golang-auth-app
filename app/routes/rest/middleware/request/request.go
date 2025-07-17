package request

import (
	loggerenum "golang-auth-app/app/common/enums/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Wrapper struct {
	Handler fiber.Handler
}

func New(logger *zap.Logger) Wrapper {
	handler := func(c *fiber.Ctx) error {
		requestId := uuid.NewString()
		c.Locals("requestId", requestId)

		zappers := []zap.Field{
			zap.String(loggerenum.RequestId.ToString(), requestId),
			zap.String(loggerenum.Path.ToString(), c.Path()),
			zap.String(loggerenum.Method.ToString(), c.Method()),
			zap.Any(loggerenum.QueryParams.ToString(), c.Queries()),
			zap.ByteString(loggerenum.RequestBody.ToString(), c.Request().Body()),
		}

		logger.Info(string(loggerenum.APIRequest), zappers...)

		return c.Next()
	}

	return Wrapper{
		Handler: handler,
	}
}
