package rest

import (
	"golang-auth-app/app/routes/rest/authentication"
	"golang-auth-app/app/routes/rest/authorization"
	"golang-auth-app/app/routes/rest/healthz"
	"golang-auth-app/app/routes/rest/management"
	"golang-auth-app/app/routes/rest/middleware"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("routes/rest",
	fx.Provide(
		initRouter,
	),

	healthz.Module,
	middleware.Module,
	authentication.Module,
	authorization.Module,
	management.Module,
)

func initRouter(app *fiber.App) fiber.Router {
	router := app.Group("/api/v1")
	return router
}
