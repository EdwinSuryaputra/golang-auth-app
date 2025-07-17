package jwt

import (
	"context"
	"encoding/json"
	"fmt"
	"golang-auth-app/config"
	"time"

	"github.com/golang-jwt/jwt"

	jwtDto "golang-auth-app/app/interfaces/authorization/jwt/dto"
)

func (i *impl) GenerateAuthToken(
	ctx context.Context,
	payload jwtDto.GenerateAuthTokenPayload,
) (*jwtDto.GenerateAuthTokenResult, error) {
	jwtSecret := config.Module.Auth.Jwt.SecretKey

	tokenDuration := config.Module.Auth.Jwt.TokenExpiryInSeconds
	if payload.IsKeepLoggedIn {
		tokenDuration = config.Module.Auth.Jwt.ForeverLoginDurationInSeconds
		if config.Module.Auth.Jwt.IsForeverLoginPersistForever {
			tokenDuration = 0
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &jwt.MapClaims{
		"username": payload.User.Username,
		"iat":      time.Now().Unix(),
	})
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}
	tokenString = fmt.Sprintf("Bearer %s", tokenString)

	redisKey := i.getRedisKeyName(payload.User.Username)
	redisVal, _ := json.Marshal(&jwtDto.AuthTokenValue{
		UserId:      payload.User.ID,
		Username:    payload.User.Username,
		FullName:    payload.User.FullName,
		Roles:       payload.Roles,
		Resources:   payload.Resources,
		TokenString: tokenString,
	})
	if err = i.genericRedisAdapter.SetValueIntoRedis(ctx, redisKey, string(redisVal), tokenDuration); err != nil {
		return nil, err
	}

	return &jwtDto.GenerateAuthTokenResult{
		TokenString: tokenString,
	}, nil
}
