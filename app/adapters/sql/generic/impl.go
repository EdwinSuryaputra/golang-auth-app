package generic

import (
	"golang-auth-app/app/interfaces/generic"

	"gorm.io/gorm"
)

type impl struct {
	db *gorm.DB
}

func New(
	db *gorm.DB,
) generic.AdapterSQL {
	return &impl{
		db: db,
	}
}
