package management

import (
	burequestbucket "golang-auth-app/app/services/management/bu_request_bucket"
	"golang-auth-app/app/services/management/role"
	"golang-auth-app/app/services/management/user"

	"go.uber.org/fx"
)

var Module = fx.Module("services/management",
	fx.Provide(
		role.New,
		user.New,
		burequestbucket.New,
	),
)
