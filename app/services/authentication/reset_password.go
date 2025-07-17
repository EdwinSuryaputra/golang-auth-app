package authentication

import (
	"context"
	"time"

	"golang-auth-app/app/datasources/sql/gorm/model"

	shautil "golang-auth-app/app/utils/sha"
)

func (i *impl) ResetPassword(ctx context.Context, tokenString, newPassword string) error {
	now := time.Now()

	redisObj, err := i.ValidateResetPasswordToken(ctx, tokenString)
	if err != nil {
		return err
	}

	encryptedNewPassword := shautil.EncryptString(newPassword)

	if err = i.userSqlAdapter.UpdateUser(ctx, &model.User{
		ID:                redisObj.UserId,
		Password:          encryptedNewPassword,
		IsDefaultPassword: false,
		UpdatedBy:         redisObj.Username,
		UpdatedAt:         now,
	}); err != nil {
		return err
	}

	if err = i.genericRedisAdapter.DeleteKeysOnRedis(ctx, []string{i.getTokenKeyName(tokenString)}); err != nil {
		return err
	}

	return nil
}
