package libraries

import (
	"golang-auth-app/app/libraries/casbin"
	"golang-auth-app/app/libraries/resty"

	"go.uber.org/fx"
)

var Module = fx.Module("libraries",
	casbin.Module,
	resty.Module,
)
