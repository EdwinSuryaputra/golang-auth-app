package middleware

import (
	"go.uber.org/fx"

	authorizationMiddleware "golang-auth-app/app/routes/rest/middleware/authorization"
	requestMiddleware "golang-auth-app/app/routes/rest/middleware/request"
	responseMiddleware "golang-auth-app/app/routes/rest/middleware/response"
)

var Module = fx.Module("routes/rest/middleware",
	fx.Provide(
		authorizationMiddleware.New,
		requestMiddleware.New,
		responseMiddleware.New,
	),
)
