package application

import (
	"context"
	"golang-auth-app/app/datasources/sql/gorm/model"
)

type AdapterSQL interface {
	GetApplicationById(ctx context.Context, applicationId int32) (*model.Application, error)

	GetApplicationByName(ctx context.Context, applicationName string) (*model.Application, error)
}
