package casbin

import (
	"github.com/casbin/casbin/v2"
	cm "github.com/casbin/casbin/v2/model"
	gormAdapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

var Module = fx.Module("libraries/casbin", fx.Provide(
	initCasbin,
))

func initCasbin(db *gorm.DB) *casbin.Enforcer {
	modelText := getCasbinModel()
	model, err := cm.NewModelFromString(modelText)
	if err != nil {
		panic(err)
	}

	adapter, err := gormAdapter.NewAdapterByDB(db)
	if err != nil {
		panic(err)
	}

	e, err := casbin.NewEnforcer(model, adapter)
	if err != nil {
		panic(err)
	}

	if err := e.LoadPolicy(); err != nil {
		panic(err)
	}

	return e
}

func getCasbinModel() string {
	return `
		[request_definition]
		r = sub, act

		[policy_definition]
		p = sub, act

		[policy_effect]
		e = some(where (p.eft == allow))

		[matchers]
		m = r.sub == p.sub && r.act == p.act
	`
}
