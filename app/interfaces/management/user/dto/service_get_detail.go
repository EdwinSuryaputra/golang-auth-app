package dto

import "time"

type ServiceGetDetailResult struct {
	Id            string                          `json:"id"`
	Username      string                          `json:"username"`
	Fullname      string                          `json:"fullname"`
	Email         string                          `json:"email"`
	Type          string                          `json:"type"`
	Status        string                          `json:"status"`
	StartDate     *time.Time                      `json:"startDate"`
	EndDate       *time.Time                      `json:"endDate"`
	AssignedRoles []*ServiceGetDetailAssignedRole `json:"assignedRoles"`
	Supplier      *ServiceGetDetailSupplier       `json:"supplier"`
	BusinessUnit  *ServiceGetDetailBusinessUnit   `json:"businessUnit"`
	CreatedAt     time.Time                       `json:"createdAt"`
	CreatedBy     string                          `json:"createdBy"`
	UpdatedAt     time.Time                       `json:"updatedAt"`
	UpdatedBy     string                          `json:"updatedBy"`
	ActivityLogId *string                         `json:"activityLogId"`
}

type ServiceGetDetailAssignedRole struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ServiceGetDetailSupplier struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ServiceGetDetailBusinessUnit struct {
	Level  string `json:"level"`
	Id     string `json:"id"`
	Name   string `json:"name"`
	Status string `json:"status"`
}
