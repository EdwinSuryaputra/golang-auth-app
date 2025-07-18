package redis

import (
	"fmt"
	"golang-auth-app/app/adapters/redis/generic"
	"golang-auth-app/config"

	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

var Module = fx.Module("adapters/redis",
	fx.Provide(
		initRedis,

		generic.New,
	),
)

func initRedis() *redis.Client {
	cfg := config.Datasource.Redis
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
	})
}
