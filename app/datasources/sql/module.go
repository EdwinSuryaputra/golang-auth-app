package sql

import (
	"golang-auth-app/app/datasources/sql/gorm"

	"go.uber.org/fx"
)

var Module = fx.Module("datasources/sql",
	gorm.Module,
)
