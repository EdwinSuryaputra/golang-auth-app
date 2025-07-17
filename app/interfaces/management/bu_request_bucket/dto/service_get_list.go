package dto

import "time"

type ServiceGetListPendingPayload struct {
	RequestDate *time.Time
	Username    string
	Fullname    string
	BuLevel     string
	BuLocation  string
	Page        int
	Limit       int
}

type ServiceGetListCompletedPayload struct {
	RequestDate *time.Time
	Username    string
	Fullname    string
	BuLevel     string
	BuLocation  string
	Status      string
	ReviewedBy  string
	ReviewedAt  *time.Time
	Page        int
	Limit       int
}

type ServiceGetListPendingResult struct {
	Entries  []*ServiceGetListPendingEntry
	TotalRow int
}

type ServiceGetListPendingEntry struct {
	Id          string     `json:"id"`
	RequestDate *time.Time `json:"requestDate"`
	UserId      string     `json:"userId"`
	Username    string     `json:"username"`
	Fullname    string     `json:"fullname"`
	BuLevel     string     `json:"buLevel"`
	BuLocation  string     `json:"buLocation"`
}

type ServiceGetListCompletedResult struct {
	Entries  []*ServiceGetListCompletedEntry
	TotalRow int
}

type ServiceGetListCompletedEntry struct {
	Id          string     `json:"id"`
	RequestDate *time.Time `json:"requestDate"`
	UserId      string     `json:"userId"`
	Username    string     `json:"username"`
	Fullname    string     `json:"fullname"`
	BuLevel     string     `json:"buLevel"`
	BuLocation  string     `json:"buLocation"`
	Status      string     `json:"status"`
	ReviewedAt  *time.Time `json:"reviewedAt"`
	ReviewedBy  *string    `json:"reviewedBy"`
}
