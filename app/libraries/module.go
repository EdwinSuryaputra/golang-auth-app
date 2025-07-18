package libraries

import (
	"golang-auth-app/app/libraries/casbin"

	"go.uber.org/fx"
)

var Module = fx.Module("libraries",
	casbin.Module,
)
