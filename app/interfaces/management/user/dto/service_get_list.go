package dto

type ServiceGetListPayload struct {
	Query *AdapterSqlGetPaginatedPayload
}

type ServiceGetListResult struct {
	Entries  []*AdapterSqlGetPaginatedEntry
	TotalRow int
}
