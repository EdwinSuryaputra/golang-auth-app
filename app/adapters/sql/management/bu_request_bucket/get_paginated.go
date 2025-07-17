package burequestbucket

import (
	"context"

	statusenum "golang-auth-app/app/common/enums/status"
	"golang-auth-app/app/datasources/sql/gorm/query"

	burbDto "golang-auth-app/app/interfaces/management/bu_request_bucket/dto"

	"gorm.io/gen/field"
	"gorm.io/gorm"
)

func (i *impl) GetPaginated(
	ctx context.Context,
	payload *burbDto.AdapterSqlGetPaginatedPayload,
) (*burbDto.AdapterSqlGetPaginatedResult, error) {
	q := query.Use(i.db.WithContext(ctx)).BuRequestBucket

	qq := q.WithContext(ctx)

	if payload.RequestDate != nil {
		requestDate := *payload.RequestDate
		qq = qq.Where(
			q.CreatedAt.Gte(requestDate),
			q.CreatedAt.Lt(requestDate.AddDate(0, 0, 1)),
		)
	}

	if len(payload.UserIds) > 0 {
		qq = qq.Where(q.UserID.In(payload.UserIds...))
	}

	if payload.BuLevel != "" {
		qq = qq.Where(q.BusinessUnitLevel.Like("%" + payload.BuLevel + "%"))
	}

	if payload.BuLocation != "" {
		qq = qq.Where(q.BusinessUnitLocation.Eq("%" + payload.BuLocation + "%"))
	}

	if payload.Status != "" {
		qq = qq.Where(q.Status.Like("%" + payload.Status + "%"))
	}

	if payload.IsCompletedOnly {
		qq = qq.Where(q.Status.In(statusenum.Rejected.ToString(), statusenum.Approved.ToString()))
	}

	if payload.ReviewedAt != nil {
		reviewedAt := *payload.ReviewedAt
		qq = qq.Where(
			q.ReviewedAt.Gte(reviewedAt),
			q.ReviewedAt.Lt(reviewedAt.AddDate(0, 0, 1)),
		)
	}

	if payload.ReviewedBy != "" {
		qq = qq.Where(q.ReviewedBy.Like("%" + payload.ReviewedBy + "%"))
	}

	qq = qq.Where(
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	)

	qq = qq.Select(
		q.ID.As("id"),
		q.CreatedAt.As("requestDate"),
		q.UserID.As("userId"),
		q.BusinessUnitLevel.As("buLevel"),
		q.BusinessUnitLocation.As("buLocation"),
		q.Status.As("status"),
		q.ReviewedAt.As("reviewedAt"),
		q.ReviewedBy.As("reviewedBy"),
	)

	f := field.Field(q.UpdatedAt)
	qq = qq.Order(f.Desc())

	offset := payload.Limit * (payload.Page - 1)

	var entries []*burbDto.AdapterSqlGetPaginatedEntry
	totalRows, err := qq.ScanByPage(&entries, offset, payload.Limit)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &burbDto.AdapterSqlGetPaginatedResult{
		Entries:  entries,
		TotalRow: int(totalRows),
	}, nil
}
