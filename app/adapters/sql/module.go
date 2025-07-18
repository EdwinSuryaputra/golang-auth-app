package sqladapters

import (
	"golang-auth-app/app/adapters/sql/activitylog"
	"golang-auth-app/app/adapters/sql/application"
	"golang-auth-app/app/adapters/sql/generic"
	"golang-auth-app/app/adapters/sql/gorm"
	"golang-auth-app/app/adapters/sql/management"
	"golang-auth-app/app/adapters/sql/resource"

	"go.uber.org/fx"
)

var Module = fx.Module("adapters/sql",
	gorm.Module,
	management.Module,

	fx.Provide(
		generic.New,
		application.New,
		resource.New,
		activitylog.New,
	),
)
