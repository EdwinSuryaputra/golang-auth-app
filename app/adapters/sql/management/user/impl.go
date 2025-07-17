package user

import (
	"golang-auth-app/app/interfaces/management/user"

	"gorm.io/gorm"
)

type impl struct {
	db *gorm.DB
}

func New(
	db *gorm.DB,
) user.AdapterSQL {
	return &impl{
		db: db,
	}
}
