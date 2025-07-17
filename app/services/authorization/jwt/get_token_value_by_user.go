package jwt

import (
	"context"
	"encoding/json"

	jwtDto "golang-auth-app/app/interfaces/authorization/jwt/dto"
)

func (i *impl) GetTokenValueByUser(ctx context.Context, username string) (*jwtDto.AuthTokenValue, error) {
	key := i.getRedisKeyName(username)

	val, err := i.genericRedisAdapter.GetValueFromRedis(ctx, key)
	if err != nil {
		return nil, err
	} else if val == "" {
		return nil, nil
	}

	var unmarshalledVal *jwtDto.AuthTokenValue
	if err = json.Unmarshal([]byte(val), &unmarshalledVal); err != nil {
		return nil, err
	}

	return unmarshalledVal, nil
}
