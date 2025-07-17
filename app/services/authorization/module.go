package authorization

import (
	"golang-auth-app/app/services/authorization/casbin"
	"golang-auth-app/app/services/authorization/jwt"

	"go.uber.org/fx"
)

var Module = fx.Module("services/authorization",
	fx.Provide(
		jwt.New,
		casbin.New,
	),
)
