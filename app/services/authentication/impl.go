package authentication

import (
	authnInterface "golang-auth-app/app/interfaces/authentication"
	jwtInterface "golang-auth-app/app/interfaces/authorization/jwt"
	genericInterface "golang-auth-app/app/interfaces/generic"
	roleInterface "golang-auth-app/app/interfaces/management/role"
	userInterface "golang-auth-app/app/interfaces/management/user"
	resourceInterface "golang-auth-app/app/interfaces/resource"
	smtpInterface "golang-auth-app/app/interfaces/smtp"
)

type impl struct {
	userSqlAdapter      userInterface.AdapterSQL
	roleSqlAdapter      roleInterface.AdapterSQL
	resourceService     resourceInterface.Service
	jwtService          jwtInterface.Service
	genericRedisAdapter genericInterface.AdapterRedis
	smtpAdapter         smtpInterface.AdapterSMTP
}

func New(
	userSqlAdapter userInterface.AdapterSQL,
	roleSqlAdapter roleInterface.AdapterSQL,
	resourceService resourceInterface.Service,
	jwtService jwtInterface.Service,
	genericRedisAdapter genericInterface.AdapterRedis,
	smtpAdapter smtpInterface.AdapterSMTP,
) authnInterface.Service {
	return &impl{
		userSqlAdapter:      userSqlAdapter,
		roleSqlAdapter:      roleSqlAdapter,
		resourceService:     resourceService,
		jwtService:          jwtService,
		genericRedisAdapter: genericRedisAdapter,
		smtpAdapter:         smtpAdapter,
	}
}
