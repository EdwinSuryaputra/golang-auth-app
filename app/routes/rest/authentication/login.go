package authentication

import (
	"fmt"
	loginmethodenum "golang-auth-app/app/common/enums/authentication/login_method"

	authenticationInterface "golang-auth-app/app/interfaces/authentication"
	authenticationDto "golang-auth-app/app/interfaces/authentication/dto"

	"golang-auth-app/app/interfaces/errorcode"

	"github.com/gofiber/fiber/v2"
)

type loginAPIPayload struct {
	Username       *string `json:"username"`
	Password       *string `json:"password"`
	LoginMethod    string  `json:"loginMethod"`
	IsForeverLogin bool    `json:"isForeverLogin"`
}

type loginAPIResponse struct {
	UserId            string `json:"userId"`
	Username          string `json:"username"`
	FullName          string `json:"fullName"`
	Token             string `json:"token"`
	IsDefaultPassword bool   `json:"isDefaultPassword"`
}

func login(
	router fiber.Router,
	authenticationService authenticationInterface.Service,
) {
	routePath := fmt.Sprintf("%s/login", prefix)
	router.Post(routePath, func(c *fiber.Ctx) error {
		ctx := c.UserContext()

		var payload *loginAPIPayload
		if err := c.BodyParser(&payload); err != nil {
			return errorcode.ErrCodeBadRequest
		}

		servicePayload, err := loginPayloadValidation(payload)
		if err != nil {
			return err
		}

		loginResult, err := authenticationService.Login(ctx, servicePayload)
		if err != nil {
			return err
		}

		return c.JSON(&loginAPIResponse{
			UserId:            loginResult.UserId,
			Username:          loginResult.Username,
			FullName:          loginResult.FullName,
			Token:             loginResult.TokenString,
			IsDefaultPassword: loginResult.IsDefaultPassword,
		})
	})
}

func loginPayloadValidation(payload *loginAPIPayload) (*authenticationDto.LoginPayload, error) {
	if payload.LoginMethod == "" {
		payload.LoginMethod = string(loginmethodenum.CredentialBased)
	}

	var servicePayload *authenticationDto.LoginPayload
	if payload.LoginMethod == string(loginmethodenum.CredentialBased) {
		if payload.Username == nil || *payload.Username == "" {
			return nil, errorcode.ErrCodeInvalidUsername
		}

		if payload.Password == nil || *payload.Password == "" {
			return nil, errorcode.ErrCodeInvalidPassword
		}

		servicePayload = &authenticationDto.LoginPayload{
			CredentialBasedPayload: &authenticationDto.CredentialBasedPayload{
				Username: *payload.Username,
				Password: *payload.Password,
			},
		}
	}

	return servicePayload, nil
}
