package burequestbucket

import (
	"context"

	"golang-auth-app/app/datasources/sql/gorm/model"
	"golang-auth-app/app/interfaces/management/bu_request_bucket/dto"
)

type AdapterSQL interface {
	GetPaginated(ctx context.Context, payload *dto.AdapterSqlGetPaginatedPayload) (*dto.AdapterSqlGetPaginatedResult, error)

	GetBuRequestBucketById(ctx context.Context, burbId int32) (*model.BuRequestBucket, error)

	InsertBuRequestBucket(ctx context.Context, payload *model.BuRequestBucket) error

	UpdateBuRequestBucket(ctx context.Context, payload *model.BuRequestBucket) error
}
