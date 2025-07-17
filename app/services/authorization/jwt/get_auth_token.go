package jwt

import (
	"context"
	"golang-auth-app/config"
	"strings"

	"github.com/golang-jwt/jwt"

	jwtDto "golang-auth-app/app/interfaces/authorization/jwt/dto"

	"golang-auth-app/app/interfaces/errorcode"
)

func (i *impl) GetAuthToken(ctx context.Context, tokenString string) (*jwtDto.AuthTokenValue, error) {
	cfg := config.Module.Auth.Jwt

	splittedToken := strings.Split(tokenString, " ")
	if splittedToken[0] != "Bearer" || splittedToken[1] == "" {
		return nil, errorcode.ErrCodeInvalidTokenFormat
	}

	token, err := jwt.Parse(splittedToken[1], func(token *jwt.Token) (any, error) {
		return []byte(cfg.SecretKey), nil
	})
	if err != nil {
		return nil, errorcode.ErrCodeInvalidTokenFormat
	} else if !token.Valid {
		return nil, nil
	}

	username := ""
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if usernameClaim, isExist := claims["username"]; isExist {
			username = usernameClaim.(string)
		} else {
			return nil, nil
		}
	} else {
		return nil, nil
	}

	authTokenValue, err := i.GetTokenValueByUser(ctx, username)
	if err != nil {
		return nil, err
	} else if authTokenValue == nil {
		return nil, errorcode.ErrCodeInvalidTokenExpired
	}

	return authTokenValue, nil
}
