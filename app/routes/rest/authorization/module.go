package authorization

import "go.uber.org/fx"

var prefix = "/authz"

var Module = fx.Module("routes/rest/authorization", fx.Invoke(
	validateResource,
	getResourceMenu,
))
