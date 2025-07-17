package dto

import (
	reviewenum "golang-auth-app/app/common/enums/review"
)

type ServiceReviewPayload struct {
	Id       int32
	Action   reviewenum.Action
	Modifier string
}
