package resource

import (
	resourceInterface "golang-auth-app/app/interfaces/resource"
)

type impl struct {
	resourceSqlAdapter resourceInterface.AdapterSQL
}

func New(
	resourceSqlAdapter resourceInterface.AdapterSQL,
) resourceInterface.Service {
	return &impl{
		resourceSqlAdapter: resourceSqlAdapter,
	}
}
