package user

import (
	activityLogInterface "golang-auth-app/app/interfaces/activity_log"
	casbinInterface "golang-auth-app/app/interfaces/authorization/casbin"
	roleInterface "golang-auth-app/app/interfaces/management/role"
	userInterface "golang-auth-app/app/interfaces/management/user"
	smtpInterface "golang-auth-app/app/interfaces/smtp"
)

type impl struct {
	roleSqlAdapter         roleInterface.AdapterSQL
	userSqlAdapter         userInterface.AdapterSQL
	casbinService          casbinInterface.Service
	smtpAdapter            smtpInterface.AdapterSMTP
	activityLogHttpAdapter activityLogInterface.AdapterSql
}

func New(
	roleSqlAdapter roleInterface.AdapterSQL,
	userSqlAdapter userInterface.AdapterSQL,
	casbinService casbinInterface.Service,
	smtpAdapter smtpInterface.AdapterSMTP,
	activityLogHttpAdapter activityLogInterface.AdapterSql,
) userInterface.Service {
	return &impl{
		roleSqlAdapter:         roleSqlAdapter,
		userSqlAdapter:         userSqlAdapter,
		casbinService:          casbinService,
		smtpAdapter:            smtpAdapter,
		activityLogHttpAdapter: activityLogHttpAdapter,
	}
}
