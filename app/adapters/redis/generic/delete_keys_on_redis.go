package generic

import "context"

func (i *impl) DeleteKeysOnRedis(ctx context.Context, keys []string) error {
	return i.redisClient.Del(ctx, keys...).Err()
}
