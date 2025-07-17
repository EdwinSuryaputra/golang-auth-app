package management

import (
	"go.uber.org/fx"

	burequestbucket "golang-auth-app/app/routes/rest/management/bu_request_bucket"
	"golang-auth-app/app/routes/rest/management/role"
	"golang-auth-app/app/routes/rest/management/user"
)

var Module = fx.Module("routes/rest/management",
	user.Module,
	role.Module,
	burequestbucket.Module,
)
