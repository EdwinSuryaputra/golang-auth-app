package authentication

import (
	"context"
	"golang-auth-app/app/common/errorcode"
	shautil "golang-auth-app/app/utils/sha"
	"time"
)

func (i *impl) ChangePassword(ctx context.Context, tokenString, oldPassword, newPassword string) error {
	now := time.Now()

	authTokenValue, err := i.jwtService.GetAuthToken(ctx, tokenString)
	if err != nil {
		return err
	} else if authTokenValue == nil {
		return errorcode.ErrCodeInvalidTokenExpired
	}

	existingUser, err := i.userSqlAdapter.GetUserById(ctx, authTokenValue.UserId)
	if err != nil {
		return err
	} else if existingUser.Password != shautil.EncryptString(oldPassword) {
		return errorcode.ErrCodeInvalidOldPassword
	}

	existingUser.Password = shautil.EncryptString(newPassword)
	existingUser.IsDefaultPassword = false
	existingUser.UpdatedAt = now
	existingUser.UpdatedBy = existingUser.Username

	if err := i.userSqlAdapter.UpdateUser(ctx, existingUser); err != nil {
		return err
	}

	return nil
}
