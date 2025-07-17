package resource

import (
	"golang-auth-app/app/interfaces/resource"

	"gorm.io/gorm"
)

type impl struct {
	db *gorm.DB
}

func New(
	db *gorm.DB,
) resource.AdapterSQL {
	return &impl{
		db: db,
	}
}
