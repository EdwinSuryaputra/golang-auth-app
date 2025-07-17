package generic

import (
	redisInterface "golang-auth-app/app/interfaces/generic"

	"github.com/redis/go-redis/v9"
)

type impl struct {
	redisClient *redis.Client
}

func New(
	redisClient *redis.Client,
) redisInterface.AdapterRedis {
	return &impl{
		redisClient: redisClient,
	}
}
