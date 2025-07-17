package adapters

import (
	httpAdapter "golang-auth-app/app/adapters/http"
	redisAdapter "golang-auth-app/app/adapters/redis"
	smtpAdapter "golang-auth-app/app/adapters/smtp"
	sqlAdapter "golang-auth-app/app/adapters/sql"

	"go.uber.org/fx"
)

var Module = fx.Module("adapters",
	sqlAdapter.Module,
	redisAdapter.Module,
	httpAdapter.Module,

	fx.Provide(smtpAdapter.New),
)
