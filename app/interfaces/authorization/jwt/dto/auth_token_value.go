package dto

import (
	resourceDto "golang-auth-app/app/interfaces/resource/dto"
)

type AuthTokenValue struct {
	UserId      int32  `json:"userId"`
	Username    string `json:"username"`
	FullName    string `json:"fullName"`
	Roles       []*AuthTokenRoleValue
	Resources   *resourceDto.ResourcePubObj `json:"resources"`
	TokenString string                      `json:"tokenString"`
}

type AuthTokenRoleValue struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}
