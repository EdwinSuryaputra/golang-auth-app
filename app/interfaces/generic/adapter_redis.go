package generic

import "context"

type AdapterRedis interface {
	GetValueFromRedis(ctx context.Context, key string) (string, error)

	SetValueIntoRedis(ctx context.Context, key, value string, duration int) error

	DeleteKeysOnRedis(ctx context.Context, keys []string) error
}
