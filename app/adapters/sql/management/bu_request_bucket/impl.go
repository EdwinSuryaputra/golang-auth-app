package burequestbucket

import (
	burequestbucket "golang-auth-app/app/interfaces/management/bu_request_bucket"

	"gorm.io/gorm"
)

type impl struct {
	db *gorm.DB
}

func New(
	db *gorm.DB,
) burequestbucket.AdapterSQL {
	return &impl{
		db: db,
	}
}
