package dto

import "time"

type AdapterSqlGetPaginatedPayload struct {
	Username     string
	Fullname     string
	Type         string
	SupplierName string
	BuLevel      string
	BuLocation   string
	Status       string
	StartDate    *time.Time
	EndDate      *time.Time
	CreatedBy    string
	CreatedAt    *time.Time
	UpdatedBy    string
	UpdatedAt    *time.Time
	Page         int
	Limit        int
}

type AdapterSqlGetPaginatedResult struct {
	Entries  []*AdapterSqlGetPaginatedEntry
	TotalRow int
}

type AdapterSqlGetPaginatedEntry struct {
	ID           int32      `json:"id" gorm:"column:id"`
	Username     string     `json:"username" gorm:"column:username"`
	FullName     string     `json:"fullname" gorm:"column:fullName"`
	Type         string     `json:"type" gorm:"column:type"`
	SupplierName *string    `json:"supplierName" gorm:"column:supplierName"`
	BuLevel      *string    `json:"buLevel" gorm:"column:buLevel"`
	BuLocation   *string    `json:"buLocation" gorm:"column:buLocation"`
	Status       string     `json:"status" gorm:"column:status"`
	StartDate    *time.Time `json:"startDate" gorm:"column:startDate"`
	EndDate      *time.Time `json:"endDate" gorm:"column:endDate"`
	CreatedBy    string     `json:"createdBy" gorm:"column:createdBy"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy    string     `json:"updatedBy" gorm:"column:updatedBy"`
	UpdatedAt    time.Time  `json:"updatedAt" gorm:"column:updatedAt"`
}
