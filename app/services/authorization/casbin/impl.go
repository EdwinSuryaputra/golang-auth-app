package casbin

import (
	casbinInterface "golang-auth-app/app/interfaces/authorization/casbin"
	roleInterface "golang-auth-app/app/interfaces/management/role"
	resourceInterface "golang-auth-app/app/interfaces/resource"

	"github.com/casbin/casbin/v2"
)

type impl struct {
	casbinEnforcer     *casbin.Enforcer
	roleSqlAdapter     roleInterface.AdapterSQL
	resourceSqlAdapter resourceInterface.AdapterSQL
}

func New(
	casbinEnforcer *casbin.Enforcer,
	roleSqlAdapter roleInterface.AdapterSQL,
	resourceSqlAdapter resourceInterface.AdapterSQL,
) casbinInterface.Service {
	return &impl{
		casbinEnforcer:     casbinEnforcer,
		roleSqlAdapter:     roleSqlAdapter,
		resourceSqlAdapter: resourceSqlAdapter,
	}
}
