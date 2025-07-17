package user

import (
	"context"

	"golang-auth-app/app/datasources/sql/gorm/query"

	userDto "golang-auth-app/app/interfaces/management/user/dto"

	"gorm.io/gen/field"
	"gorm.io/gorm"
)

func (i *impl) GetPaginated(
	ctx context.Context,
	payload *userDto.AdapterSqlGetPaginatedPayload,
) (*userDto.AdapterSqlGetPaginatedResult, error) {
	userQuery := query.Use(i.db.WithContext(ctx)).User

	qq := userQuery.WithContext(ctx)

	if payload.Username != "" {
		qq = qq.Where(userQuery.Username.Like("%" + payload.Username + "%"))
	}

	if payload.Fullname != "" {
		qq = qq.Where(userQuery.FullName.Like("%" + payload.Fullname + "%"))
	}

	if payload.Type != "" {
		qq = qq.Where(userQuery.Type.Like("%" + payload.Type + "%"))
	}

	if payload.SupplierName != "" {
		qq = qq.Where(userQuery.Type.Like("%" + payload.SupplierName + "%"))
	}

	if payload.BuLevel != "" {
		qq = qq.Where(userQuery.Type.Like("%" + payload.BuLevel + "%"))
	}

	if payload.BuLocation != "" {
		qq = qq.Where(userQuery.Type.Like("%" + payload.BuLocation + "%"))
	}

	if payload.Status != "" {
		qq = qq.Where(userQuery.Status.Like("%" + payload.Status + "%"))
	}

	if payload.StartDate != nil {
		startDate := *payload.StartDate
		qq = qq.Where(
			userQuery.StartDate.Gte(startDate),
			userQuery.StartDate.Lt(startDate.AddDate(0, 0, 1)),
		)
	}

	if payload.EndDate != nil {
		endDate := *payload.EndDate
		qq = qq.Where(
			userQuery.EndDate.Gte(endDate),
			userQuery.EndDate.Lt(endDate.AddDate(0, 0, 1)),
		)
	}

	if payload.CreatedBy != "" {
		qq = qq.Where(userQuery.CreatedBy.Like("%" + payload.CreatedBy + "%"))
	}

	if payload.CreatedAt != nil {
		createdAt := *payload.CreatedAt
		qq = qq.Where(
			userQuery.CreatedAt.Gte(createdAt),
			userQuery.CreatedAt.Lt(createdAt.AddDate(0, 0, 1)),
		)
	}

	if payload.UpdatedBy != "" {
		qq = qq.Where(userQuery.UpdatedBy.Like("%" + payload.UpdatedBy + "%"))
	}

	if payload.UpdatedAt != nil {
		updatedAt := *payload.UpdatedAt
		qq = qq.Where(
			userQuery.UpdatedAt.Gte(updatedAt),
			userQuery.UpdatedAt.Lt(updatedAt.AddDate(0, 0, 1)),
		)
	}

	qq = qq.Select(
		userQuery.ID.As("id"),
		userQuery.Username.As("username"),
		userQuery.FullName.As("fullName"),
		userQuery.Type.As("type"),
		userQuery.SupplierName.As("supplierName"),
		userQuery.BusinessUnitLevel.As("buLevel"),
		userQuery.BusinessUnitLocation.As("buLocation"),
		userQuery.Status.As("status"),
		userQuery.StartDate.As("startDate"),
		userQuery.EndDate.As("endDate"),
		userQuery.CreatedAt.As("createdAt"),
		userQuery.CreatedBy.As("createdBy"),
		userQuery.UpdatedAt.As("updatedAt"),
		userQuery.UpdatedBy.As("updatedBy"),
	)

	qq = qq.Where(
		userQuery.DeletedAt.IsNull(),
		userQuery.DeletedBy.IsNull(),
	)

	f := field.Field(userQuery.UpdatedAt)
	qq = qq.Order(f.Desc())

	offset := payload.Limit * (payload.Page - 1)

	var entries []*userDto.AdapterSqlGetPaginatedEntry
	totalRows, err := qq.ScanByPage(&entries, offset, payload.Limit)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &userDto.AdapterSqlGetPaginatedResult{
		Entries:  entries,
		TotalRow: int(totalRows),
	}, nil
}
