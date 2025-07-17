package generic

import (
	"context"
	"time"
)

func (i *impl) SetValueIntoRedis(ctx context.Context, key, value string, duration int) error {
	return i.redisClient.Set(ctx, key, value, time.Duration(duration)*time.Second).Err()
}
