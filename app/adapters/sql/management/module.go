package management

import (
	burequestbucket "golang-auth-app/app/adapters/sql/management/bu_request_bucket"
	"golang-auth-app/app/adapters/sql/management/role"
	"golang-auth-app/app/adapters/sql/management/user"

	"go.uber.org/fx"
)

var Module = fx.Module("adapters/sql",
	fx.Provide(
		user.New,
		role.New,
		burequestbucket.New,
	),
)
