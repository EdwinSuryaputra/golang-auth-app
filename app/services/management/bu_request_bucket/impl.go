package burequestbucket

import (
	activityLogInterface "golang-auth-app/app/interfaces/activity_log"
	buInterface "golang-auth-app/app/interfaces/business_unit"
	burbInterface "golang-auth-app/app/interfaces/management/bu_request_bucket"
	userInterface "golang-auth-app/app/interfaces/management/user"
)

type impl struct {
	burbSqlAdapter         burbInterface.AdapterSQL
	userSqlAdapter         userInterface.AdapterSQL
	buHttpAdapter          buInterface.AdapterHttp
	activityLogHttpAdapter activityLogInterface.AdapterHttp
}

func New(
	burbSqlAdapter burbInterface.AdapterSQL,
	userSqlAdapter userInterface.AdapterSQL,
	buHttpAdapter buInterface.AdapterHttp,
	activityLogHttpAdapter activityLogInterface.AdapterHttp,
) burbInterface.Service {
	return &impl{
		burbSqlAdapter:         burbSqlAdapter,
		userSqlAdapter:         userSqlAdapter,
		buHttpAdapter:          buHttpAdapter,
		activityLogHttpAdapter: activityLogHttpAdapter,
	}
}
