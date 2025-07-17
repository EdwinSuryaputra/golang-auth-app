package application

import (
	"golang-auth-app/app/interfaces/application"

	"gorm.io/gorm"
)

type impl struct {
	db *gorm.DB
}

func New(
	db *gorm.DB,
) application.AdapterSQL {
	return &impl{
		db: db,
	}
}
