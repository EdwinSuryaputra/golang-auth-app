package role

import (
	"context"

	"golang-auth-app/app/adapters/sql/gorm/model"
	"golang-auth-app/app/adapters/sql/gorm/query"

	roleInterface "golang-auth-app/app/interfaces/management/role"

	"gorm.io/gen/field"
	"gorm.io/gorm"
)

func (i *impl) GetPaginated(
	ctx context.Context,
	payload *roleInterface.AdapterSqlGetPaginatedPayload,
) (*roleInterface.AdapterSqlGetPaginatedResult, error) {
	q := query.Use(i.db.WithContext(ctx)).Role

	qq := q.WithContext(ctx)

	if payload.Query.Name != "" {
		qq = qq.Where(q.Name.Like(payload.Query.Name + "%"))
	}

	if payload.Query.Status != "" {
		qq = qq.Where(q.Status.Like(payload.Query.Status + "%"))
	}

	if payload.Query.Description != "" {
		qq = qq.Where(q.Description.Like(payload.Query.Description + "%"))
	}

	if payload.Query.Type != "" {
		qq = qq.Where(q.Type.Like(payload.Query.Type + "%"))
	}

	if payload.Query.InactiveDate != nil {
		inactiveDate := *payload.Query.InactiveDate
		qq = qq.Where(
			q.InactiveDate.Gte(inactiveDate),
			q.InactiveDate.Lt(inactiveDate.AddDate(0, 0, 1)),
		)
	}

	if payload.Query.CreatedBy != "" {
		qq = qq.Where(q.CreatedBy.Like(payload.Query.CreatedBy + "%"))
	}

	if payload.Query.CreatedAt != nil {
		createdAt := *payload.Query.CreatedAt
		qq = qq.Where(
			q.CreatedAt.Gte(createdAt),
			q.CreatedAt.Lt(createdAt.AddDate(0, 0, 1)),
		)
	}

	if payload.Query.UpdatedBy != "" {
		qq = qq.Where(q.UpdatedBy.Like(payload.Query.UpdatedBy + "%"))
	}

	if payload.Query.UpdatedAt != nil {
		updatedAt := *payload.Query.UpdatedAt
		qq = qq.Where(
			q.UpdatedAt.Gte(updatedAt),
			q.UpdatedAt.Lt(updatedAt.AddDate(0, 0, 1)),
		)
	}

	qq = qq.Where(
		q.DeletedAt.IsNull(),
		q.DeletedBy.IsNull(),
	)

	f := field.Field(q.CreatedAt)
	qq = qq.Order(f.Desc())

	offset := payload.Limit * (payload.Page - 1)

	var entries []*model.Role
	totalRows, err := qq.ScanByPage(&entries, offset, payload.Limit)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &roleInterface.AdapterSqlGetPaginatedResult{
		Entries:  entries,
		TotalRow: int(totalRows),
	}, nil
}
