package dto

import (
	"time"
)

type ServiceGetDetailResult struct {
	Id            string                      `json:"id"`
	Name          string                      `json:"name"`
	Description   string                      `json:"description"`
	Type          string                      `json:"type"`
	Status        string                      `json:"status"`
	InactiveDate  *time.Time                  `json:"inactiveDate"`
	Resources     []*ServiceGetDetailResource `json:"resources"`
	CreatedAt     time.Time                   `json:"createdAt"`
	CreatedBy     string                      `json:"createdBy"`
	UpdatedAt     time.Time                   `json:"updatedAt"`
	UpdatedBy     string                      `json:"updatedBy"`
	ActivityLogId *string                     `json:"activityLogId"`
}

type ServiceGetDetailResource struct {
	MenuId      string                                  `json:"menuId"`
	MenuName    string                                  `json:"menuName"`
	SubmenuId   string                                  `json:"submenuId"`
	SubmenuName string                                  `json:"submenuName"`
	Functions   []*ServiceGetDetailRoleResourceFunction `json:"functions"`
}

type ServiceGetDetailRoleResourceFunction struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
