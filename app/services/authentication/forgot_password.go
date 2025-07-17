package authentication

import (
	"context"
	"encoding/json"
	"fmt"

	authDto "golang-auth-app/app/interfaces/authentication/dto"
	"golang-auth-app/app/interfaces/smtp"

	fileutil "golang-auth-app/app/utils/file"
	shautil "golang-auth-app/app/utils/sha"

	config "golang-auth-app/config"

	"github.com/google/uuid"
)

func (i *impl) ForgotPassword(ctx context.Context, email string) error {
	existingUser, err := i.userSqlAdapter.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	tokenString := shautil.EncryptString(uuid.NewString())
	redisVal, _ := json.Marshal(&authDto.ForgotPasswordRedisObj{
		UserId:      existingUser.ID,
		Username:    existingUser.Username,
		TokenString: tokenString,
	})
	if err = i.genericRedisAdapter.SetValueIntoRedis(ctx, i.getTokenKeyName(tokenString), string(redisVal), 900); err != nil {
		return err
	}

	if err = i.sendForgotPasswordEmail(ctx, email, existingUser.Username, tokenString); err != nil {
		return err
	}

	return nil
}

func (i *impl) getTokenKeyName(tokenString string) string {
	return fmt.Sprintf("forgot-password-%s", tokenString)
}

func (i *impl) sendForgotPasswordEmail(ctx context.Context, email, username, tokenString string) error {
	resetPageUrl := fmt.Sprintf("%s?token=%s", config.Module.Auth.WebPageUrl.ResetPasswordUrl, tokenString)

	type htmlTemplate struct {
		Username string
		ResetUrl string
	}
	htmlBody, err := fileutil.ParseTemplateFile(ctx, "./etc/template/authentication/forgot-password-prompt-email.html", htmlTemplate{
		Username: username,
		ResetUrl: resetPageUrl,
	})
	if err != nil {
		return err
	}

	cfgSMTP := config.ExternalService.Smtp
	if err = i.smtpAdapter.SendEmail(ctx, &smtp.SendEmailPayload{
		Host:       cfgSMTP.Host,
		Port:       cfgSMTP.Port,
		User:       cfgSMTP.User,
		Password:   cfgSMTP.Password,
		Recipients: []string{email},
		Subject:    "Reset Password Request",
		HTMLBody:   htmlBody,
	}); err != nil {
		return err
	}

	return nil
}
