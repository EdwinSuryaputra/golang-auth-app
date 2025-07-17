package smtpadapters

import (
	"golang-auth-app/app/interfaces/smtp"

	"gorm.io/gorm"
)

type impl struct {
	db *gorm.DB
}

func New(
	db *gorm.DB,
) smtp.AdapterSMTP {
	return &impl{
		db: db,
	}
}
