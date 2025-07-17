package jwt

import (
	jwtInterface "golang-auth-app/app/interfaces/authorization/jwt"
	genericInterface "golang-auth-app/app/interfaces/generic"
)

type impl struct {
	genericRedisAdapter genericInterface.AdapterRedis
}

func New(
	genericRedisAdapter genericInterface.AdapterRedis,
) jwtInterface.Service {
	return &impl{
		genericRedisAdapter: genericRedisAdapter,
	}
}
