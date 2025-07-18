package management

import (
	"golang-auth-app/app/adapters/sql/management/role"
	"golang-auth-app/app/adapters/sql/management/user"

	"go.uber.org/fx"
)

var Module = fx.Module("adapters/sql",
	fx.Provide(
		user.New,
		role.New,
	),
)
