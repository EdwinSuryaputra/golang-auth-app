package role

import (
	"golang-auth-app/app/interfaces/management/role"

	"gorm.io/gorm"
)

type impl struct {
	db *gorm.DB
}

func New(
	db *gorm.DB,
) role.AdapterSQL {
	return &impl{
		db: db,
	}
}
