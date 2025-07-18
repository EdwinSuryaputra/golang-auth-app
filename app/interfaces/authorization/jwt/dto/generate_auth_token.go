package dto

import (
	"golang-auth-app/app/adapters/sql/gorm/model"
	resourceDto "golang-auth-app/app/interfaces/resource/dto"
)

type GenerateAuthTokenPayload struct {
	User           *model.User
	Roles          []*AuthTokenRoleValue
	Resources      *resourceDto.ResourcePubObj
	IsKeepLoggedIn bool
}

type GenerateAuthTokenResult struct {
	TokenString string
}
