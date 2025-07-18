package authentication

import (
	"context"
	"encoding/json"

	"golang-auth-app/app/common/errorcode"
	authnDto "golang-auth-app/app/interfaces/authentication/dto"
)

func (i *impl) ValidateResetPasswordToken(
	ctx context.Context,
	tokenString string,
) (*authnDto.ForgotPasswordRedisObj, error) {
	getRedisResult, err := i.genericRedisAdapter.GetValueFromRedis(ctx, i.getTokenKeyName(tokenString))
	if err != nil {
		return nil, err
	} else if getRedisResult == "" {
		return nil, errorcode.WithCustomMessage(errorcode.ErrCodeUnauthorized, "Token is not exist or already expired")
	}

	var redisObj *authnDto.ForgotPasswordRedisObj
	if err = json.Unmarshal([]byte(getRedisResult), &redisObj); err != nil {
		return nil, err
	} else if redisObj == nil {
		return nil, errorcode.WithCustomMessage(errorcode.ErrCodeInternalServerError, "Redis object is empty")
	}

	return redisObj, nil
}
