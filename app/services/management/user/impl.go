package user

import (
	activityLogInterface "golang-auth-app/app/interfaces/activity_log"
	casbinInterface "golang-auth-app/app/interfaces/authorization/casbin"
	buInterface "golang-auth-app/app/interfaces/business_unit"
	burbInterface "golang-auth-app/app/interfaces/management/bu_request_bucket"
	roleInterface "golang-auth-app/app/interfaces/management/role"
	userInterface "golang-auth-app/app/interfaces/management/user"
	smtpInterface "golang-auth-app/app/interfaces/smtp"
	supplierInterface "golang-auth-app/app/interfaces/supplier"
)

type impl struct {
	roleSqlAdapter         roleInterface.AdapterSQL
	userSqlAdapter         userInterface.AdapterSQL
	buHttpAdapter          buInterface.AdapterHttp
	supplierHttpAdapter    supplierInterface.AdapterHttp
	casbinService          casbinInterface.Service
	burbSqlAdapter         burbInterface.AdapterSQL
	smtpAdapter            smtpInterface.AdapterSMTP
	activityLogHttpAdapter activityLogInterface.AdapterHttp
}

func New(
	roleSqlAdapter roleInterface.AdapterSQL,
	userSqlAdapter userInterface.AdapterSQL,
	buHttpAdapter buInterface.AdapterHttp,
	supplierHttpAdapter supplierInterface.AdapterHttp,
	casbinService casbinInterface.Service,
	burbSqlAdapter burbInterface.AdapterSQL,
	smtpAdapter smtpInterface.AdapterSMTP,
	activityLogHttpAdapter activityLogInterface.AdapterHttp,
) userInterface.Service {
	return &impl{
		roleSqlAdapter:         roleSqlAdapter,
		userSqlAdapter:         userSqlAdapter,
		buHttpAdapter:          buHttpAdapter,
		supplierHttpAdapter:    supplierHttpAdapter,
		casbinService:          casbinService,
		burbSqlAdapter:         burbSqlAdapter,
		smtpAdapter:            smtpAdapter,
		activityLogHttpAdapter: activityLogHttpAdapter,
	}
}
