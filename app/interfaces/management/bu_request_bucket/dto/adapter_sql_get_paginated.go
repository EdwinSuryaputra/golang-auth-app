package dto

import "time"

type AdapterSqlGetPaginatedPayload struct {
	RequestDate     *time.Time
	UserIds         []int32
	BuLevel         string
	BuLocation      string
	Status          string
	IsCompletedOnly bool
	ReviewedBy      string
	ReviewedAt      *time.Time
	Page            int
	Limit           int
}

type AdapterSqlGetPaginatedResult struct {
	Entries  []*AdapterSqlGetPaginatedEntry
	TotalRow int
}

type AdapterSqlGetPaginatedEntry struct {
	Id          int32      `json:"id" gorm:"column:id"`
	RequestDate *time.Time `json:"requestDate" gorm:"column:requestDate"`
	UserId      int32      `json:"userId" gorm:"column:userId"`
	BuLevel     string     `json:"buLevel" gorm:"column:buLevel"`
	BuLocation  string     `json:"buLocation" gorm:"column:buLocation"`
	Status      string     `json:"status" gorm:"column:status"`
	ReviewedBy  *string    `json:"reviewedBy" gorm:"column:reviewedBy"`
	ReviewedAt  *time.Time `json:"reviewedAt" gorm:"column:reviewedAt"`
}
