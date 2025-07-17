package dto

import "golang-auth-app/app/datasources/sql/gorm/model"

type ServiceUpdatePayload struct {
	CurrentDataUser *model.User
	NewDataUser     *model.User
}
