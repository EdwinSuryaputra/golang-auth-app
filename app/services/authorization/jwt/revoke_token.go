package jwt

import (
	"context"
)

func (i *impl) RevokeToken(ctx context.Context, token string) error {
	authTokenValue, err := i.GetAuthToken(ctx, token)
	if err != nil {
		return err
	} else if authTokenValue == nil {
		return nil
	}

	redisKey := i.getRedisKeyName(authTokenValue.Username)
	err = i.genericRedisAdapter.DeleteKeysOnRedis(ctx, []string{redisKey})
	if err != nil {
		return err
	}

	return nil
}
