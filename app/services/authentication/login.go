package authentication

import (
	"context"

	"golang-auth-app/app/adapters/sql/gorm/model"
	statusenum "golang-auth-app/app/common/enums/status"

	authenticationDto "golang-auth-app/app/interfaces/authentication/dto"
	authorizationDto "golang-auth-app/app/interfaces/authorization/jwt/dto"
	roleDto "golang-auth-app/app/interfaces/management/role/dto"
	resourceDto "golang-auth-app/app/interfaces/resource/dto"

	publicfacingutil "golang-auth-app/app/utils/publicfacing"
	shautil "golang-auth-app/app/utils/sha"
	sliceutil "golang-auth-app/app/utils/slice"

	"golang-auth-app/app/common/errorcode"
)

func (i *impl) Login(
	ctx context.Context,
	payload *authenticationDto.LoginPayload,
) (*authenticationDto.LoginResult, error) {
	var err error
	var existingUser *model.User

	if payload.CredentialBasedPayload != nil {
		existingUser, err = i.userSqlAdapter.GetUserByUsernameForLogin(ctx, payload.CredentialBasedPayload.Username)
		if err != nil {
			return nil, err
		} else if existingUser == nil {
			return nil, errorcode.ErrCodeUserNotFound
		} else if existingUser.Status != statusenum.Active.ToString() {
			return nil, errorcode.ErrCodeUserNotActive
		} else if existingUser.Password != shautil.EncryptString(payload.CredentialBasedPayload.Password) {
			return nil, errorcode.ErrCodeInvalidPassword
		}
	}

	roles, err := i.roleSqlAdapter.GetRolesByUserId(ctx, &roleDto.AdapterGetRolesByUserIdPayload{
		UserId:       existingUser.ID,
		IsActiveOnly: true,
	})
	if err != nil {
		return nil, err
	} else if len(roles) == 0 {
		return nil, errorcode.ErrCodeUserHasNoRole
	}
	roleIds := sliceutil.Map(roles, func(dt *model.Role) int32 { return dt.ID })
	applicationIds := sliceutil.Map(roles, func(dt *model.Role) int32 { return dt.ApplicationID })

	roleResources, err := i.roleSqlAdapter.GetRoleResourceMappings(ctx, roleIds)
	if err != nil {
		return nil, err
	}
	resourceIds := sliceutil.Map(roleResources, func(dt *model.RoleResourceMapping) int32 { return dt.ResourceID })

	resources, err := i.resourceService.GetPublicResource(ctx, &resourceDto.GetPublicResourcePayload{
		ApplicationIds: applicationIds,
		ResourceIds:    resourceIds,
	})
	if err != nil {
		return nil, err
	}

	authTokenValue, err := i.jwtService.GenerateAuthToken(ctx, authorizationDto.GenerateAuthTokenPayload{
		User: existingUser,
		Roles: sliceutil.Map(roles, func(dt *model.Role) *authorizationDto.AuthTokenRoleValue {
			return &authorizationDto.AuthTokenRoleValue{
				Id:   dt.ID,
				Name: dt.Name,
			}
		}),
		Resources:      resources,
		IsKeepLoggedIn: payload.IsKeepLoggedIn,
	})
	if err != nil {
		return nil, err
	}

	encodedUserId, err := publicfacingutil.Encode(existingUser.ID)
	if err != nil {
		return nil, err
	}

	return &authenticationDto.LoginResult{
		UserId:            encodedUserId,
		Username:          existingUser.Username,
		FullName:          existingUser.FullName,
		TokenString:       authTokenValue.TokenString,
		IsDefaultPassword: existingUser.IsDefaultPassword,
	}, nil
}
