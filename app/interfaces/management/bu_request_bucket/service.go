package burequestbucket

import (
	"context"
	"golang-auth-app/app/interfaces/management/bu_request_bucket/dto"
)

type Service interface {
	GetListPending(ctx context.Context, payload *dto.ServiceGetListPendingPayload) (*dto.ServiceGetListPendingResult, error)

	GetListCompleted(ctx context.Context, payload *dto.ServiceGetListCompletedPayload) (*dto.ServiceGetListCompletedResult, error)

	Review(ctx context.Context, payload *dto.ServiceReviewPayload) error
}
