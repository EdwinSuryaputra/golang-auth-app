package datasources

import (
	"golang-auth-app/app/datasources/redis"
	"golang-auth-app/app/datasources/sql"

	"go.uber.org/fx"
)

var Module = fx.Module("datasources",
	sql.Module,
	redis.Module,
)
