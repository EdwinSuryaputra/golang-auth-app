package role

import (
	reviewenum "golang-auth-app/app/common/enums/review"

	"golang-auth-app/app/adapters/sql/gorm/model"
	"time"
)

type AdapterGetRolesByNamesPayload struct {
	RoleNames  []string
	ActiveOnly bool
}

type AdapterSqlGetPaginatedPayload struct {
	Query *GetPaginatedColumnsFilter
	Page  int
	Limit int
}

type GetPaginatedColumnsFilter struct {
	Name         string
	Status       string
	Description  string
	Type         string
	InactiveDate *time.Time
	CreatedBy    string
	CreatedAt    *time.Time
	UpdatedBy    string
	UpdatedAt    *time.Time
}

type AdapterSqlGetPaginatedResult struct {
	Entries  []*model.Role
	TotalRow int
}

type ServiceGetListRoleResult struct {
	Entries  []*model.Role
	TotalRow int
}

type ServiceUpdateRolePayload struct {
	CurrentDataRole *model.Role
	NewDataRole     *model.Role
}

type ServiceReviewRolePayload struct {
	ExistingRole *model.Role
	Action       reviewenum.Action
	Modifier     string
}
