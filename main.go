package main

import (
	"context"
	"fmt"
	"golang-auth-app/app/adapters"
	"golang-auth-app/app/interfaces/authorization/casbin"
	"golang-auth-app/app/libraries"
	"golang-auth-app/app/routes/rest"
	"golang-auth-app/app/routes/rest/middleware/request"
	"golang-auth-app/app/routes/rest/middleware/response"
	"golang-auth-app/app/services"
	config "golang-auth-app/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	// Load the env config
	config.Init()

	// Force UTC timezone globally
	time.Local = time.UTC

	// Initiate the dependency injection
	app := fx.New(
		fx.Provide(
			initLogger,
			initFiber,
		),

		rest.Module,
		adapters.Module,
		libraries.Module,
		services.Module,

		fx.Invoke(
			runFiber,
			initLoadCasbinPolicies,
		),
	)

	app.Run()
}

func initFiber(requestMiddlware request.Wrapper, responseMiddlware response.Wrapper) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler:          nil,
		Prefork:               false,
		BodyLimit:             10 * 1024 * 1024,
		DisableStartupMessage: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     config.Application.Cors.Origins,
		AllowMethods:     config.Application.Cors.Methods,
		AllowHeaders:     config.Application.Cors.Headers,
		AllowCredentials: config.Application.Cors.AllowCredentials,
	}))

	app.Use(requestMiddlware.Handler)
	app.Use(responseMiddlware.Handler)

	return app
}

func runFiber(lc fx.Lifecycle, app *fiber.App, logger *zap.Logger) {
	appConfig := config.Application

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if appConfig.Tls.IsEnabled {
					logger.Info(fmt.Sprintf("🚀🛡️ %s server started on: %d", appConfig.Name, appConfig.Port))
					logger.Info("TLS mode enabled")
					if err := app.ListenTLS(fmt.Sprintf(":%d", appConfig.Port), appConfig.Tls.CertFile, appConfig.Tls.KeyFile); err != nil {
						logger.Error("", zap.Error(err))
					}
				} else {
					logger.Info(fmt.Sprintf("🚀🛡️ %s server started on: %d", appConfig.Name, appConfig.Port))
					logger.Info("TLS mode disabled")
					if err := app.Listen(fmt.Sprintf(":%d", appConfig.Port)); err != nil {
						logger.Error("", zap.Error(err))
					}
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("Shutting down Fiber...")
			return app.Shutdown()
		},
	})
}

func initLogger() *zap.Logger {
	cfg := zap.NewDevelopmentConfig()
	cfg.Encoding = "json"

	switch config.Application.Env {
	case config.Local:
		cfg.Encoding = "console"
		cfg.EncoderConfig.ConsoleSeparator = ` | `
	case config.Prod:
		cfg = zap.NewProductionConfig()
	}

	cfg.DisableStacktrace = true // Disabled due to it shows where the logger.Error is being called instead of where the error is
	cfg.EncoderConfig.CallerKey = "caller"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.TimeKey = "ts"
	cfg.EncoderConfig.NameKey = "logger"
	cfg.EncoderConfig.MessageKey = "msg"
	cfg.EncoderConfig.StacktraceKey = "stacktrace"
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}

	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	return logger
}

func initLoadCasbinPolicies(logger *zap.Logger, casbinService casbin.Service) {
	logger.Info("Loading casbin policies...")

	if err := casbinService.SyncAllPolicies(context.Background()); err != nil {
		logger.Error("Error occurred during load casbin policies", zap.Error(err))
		return
	}

	logger.Info("Casbin policies have been successfully loaded")
}
