package management

import (
	"go.uber.org/fx"

	"golang-auth-app/app/routes/rest/management/role"
	"golang-auth-app/app/routes/rest/management/user"
)

var Module = fx.Module("routes/rest/management",
	user.Module,
	role.Module,
)
