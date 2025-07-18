package activitylog

import (
	activitylog "golang-auth-app/app/interfaces/activity_log"

	"gorm.io/gorm"
)

type impl struct {
	db *gorm.DB
}

func New(
	db *gorm.DB,
) activitylog.AdapterSQL {
	return &impl{
		db: db,
	}
}
