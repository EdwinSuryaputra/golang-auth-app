package generic

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func (i *impl) GetValueFromRedis(ctx context.Context, key string) (string, error) {
	resp, err := i.redisClient.Get(ctx, key).Result()
	if err == redis.Nil || resp == "" {
		return "", nil
	} else if err != nil {
		return "", err
	}

	return resp, nil
}
