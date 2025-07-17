package redis

import (
	"golang-auth-app/app/adapters/redis/generic"

	"go.uber.org/fx"
)

var Module = fx.Module("adapters/redis",
	fx.Provide(
		generic.New,
	),
)
