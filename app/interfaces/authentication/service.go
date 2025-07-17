package authentication

import (
	"context"

	authnDto "golang-auth-app/app/interfaces/authentication/dto"
)

type Service interface {
	Login(ctx context.Context, payload *authnDto.LoginPayload) (*authnDto.LoginResult, error)

	Logout(ctx context.Context, tokenString string) error

	ForgotPassword(ctx context.Context, email string) error

	ResetPassword(ctx context.Context, tokenString, newPassword string) error

	ValidateResetPasswordToken(ctx context.Context, tokenString string) (*authnDto.ForgotPasswordRedisObj, error)

	ChangePassword(ctx context.Context, tokenString, oldPassword, newPassword string) error
}
