package role

import (
	activityLogInterface "golang-auth-app/app/interfaces/activity_log"
	casbinInterface "golang-auth-app/app/interfaces/authorization/casbin"
	jwtInterface "golang-auth-app/app/interfaces/authorization/jwt"
	roleInterface "golang-auth-app/app/interfaces/management/role"
	userInterface "golang-auth-app/app/interfaces/management/user"
	resourceInterface "golang-auth-app/app/interfaces/resource"
)

type impl struct {
	roleSqlAdapter         roleInterface.AdapterSQL
	userSqlAdapter         userInterface.AdapterSQL
	resourceSqlAdapter     resourceInterface.AdapterSQL
	casbinService          casbinInterface.Service
	jwtService             jwtInterface.Service
	activityLogHttpAdapter activityLogInterface.AdapterSql
}

func New(
	roleSqlAdapter roleInterface.AdapterSQL,
	userSqlAdapter userInterface.AdapterSQL,
	resourceSqlAdapter resourceInterface.AdapterSQL,
	casbinService casbinInterface.Service,
	jwtService jwtInterface.Service,
	activityLogHttpAdapter activityLogInterface.AdapterSql,
) roleInterface.Service {
	return &impl{
		roleSqlAdapter:         roleSqlAdapter,
		userSqlAdapter:         userSqlAdapter,
		resourceSqlAdapter:     resourceSqlAdapter,
		casbinService:          casbinService,
		jwtService:             jwtService,
		activityLogHttpAdapter: activityLogHttpAdapter,
	}
}
