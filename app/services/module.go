package services

import (
	"golang-auth-app/app/services/authentication"
	"golang-auth-app/app/services/authorization"
	"golang-auth-app/app/services/management"
	"golang-auth-app/app/services/resource"

	"go.uber.org/fx"
)

var Module = fx.Module("services",
	authorization.Module,
	management.Module,

	fx.Provide(
		authentication.New,
		resource.New,
	),
)
