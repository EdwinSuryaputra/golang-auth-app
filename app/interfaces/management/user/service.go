package user

import (
	"context"
	"golang-auth-app/app/adapters/sql/gorm/model"

	userDto "golang-auth-app/app/interfaces/management/user/dto"
)

type Service interface {
	GetList(ctx context.Context, payload *userDto.ServiceGetListPayload) (*userDto.ServiceGetListResult, error)

	GetDetail(ctx context.Context, encodedUserId string) (*userDto.ServiceGetDetailResult, error)

	CreateDraft(ctx context.Context, newUser *model.User) error

	Update(ctx context.Context, payload *userDto.ServiceUpdatePayload) error

	Review(ctx context.Context, payload *userDto.ServiceReviewPayload) error
}
