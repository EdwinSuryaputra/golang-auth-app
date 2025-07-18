package management

import (
	"golang-auth-app/app/services/management/role"
	"golang-auth-app/app/services/management/user"

	"go.uber.org/fx"
)

var Module = fx.Module("services/management",
	fx.Provide(
		role.New,
		user.New,
	),
)
