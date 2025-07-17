package dto

import (
	reviewenum "golang-auth-app/app/common/enums/review"
	"golang-auth-app/app/datasources/sql/gorm/model"
)

type ServiceReviewPayload struct {
	ExistingUser *model.User
	Action       reviewenum.Action
	Modifier     string
}
