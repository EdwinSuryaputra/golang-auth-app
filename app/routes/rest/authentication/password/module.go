package password

import (
	"go.uber.org/fx"
)

var prefix = "/authn/password"

var Module = fx.Module("routes/rest/authentication/password", fx.Invoke(
	forgotPassword,
	resetPassword,
	validateResetPasswordToken,
	changePassword,
))
