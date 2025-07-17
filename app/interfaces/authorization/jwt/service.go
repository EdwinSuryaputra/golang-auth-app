package jwt

import (
	"context"

	jwtDto "golang-auth-app/app/interfaces/authorization/jwt/dto"
)

type Service interface {
	GenerateAuthToken(
		ctx context.Context,
		payload jwtDto.GenerateAuthTokenPayload,
	) (*jwtDto.GenerateAuthTokenResult, error)

	GetAuthToken(ctx context.Context, tokenString string) (*jwtDto.AuthTokenValue, error)

	GetTokenValueByUser(ctx context.Context, username string) (*jwtDto.AuthTokenValue, error)

	RevokeToken(ctx context.Context, token string) error
}
