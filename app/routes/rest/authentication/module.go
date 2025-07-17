package authentication

import (
	"golang-auth-app/app/routes/rest/authentication/password"

	"go.uber.org/fx"
)

var prefix = "/authn"

var Module = fx.Module("routes/rest/authentication",
	password.Module,

	fx.Invoke(
		login,
		logout,
	))
