package redis

import (
	"fmt"
	config "golang-auth-app/config"

	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

var Module = fx.Module("datasources/redis",
	fx.Provide(
		InitRedis,
	),
)

func InitRedis() *redis.Client {
	cfg := config.Datasource.Redis
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
	})
}
