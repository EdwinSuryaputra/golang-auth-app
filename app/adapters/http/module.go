package httpadapters

import (
	activitylog "golang-auth-app/app/adapters/http/activity_log"
	businessunit "golang-auth-app/app/adapters/http/business_unit"
	"golang-auth-app/app/adapters/http/supplier"

	"go.uber.org/fx"
)

var Module = fx.Module("adapters/http",
	fx.Provide(
		businessunit.New,
		supplier.New,
		activitylog.New,
	),
)
