package request

type GetListAPIRequest[T any] struct {
	Page           int `json:"page"`
	Limit          int `json:"limit"`
	PropertyFilter *T  `json:"propertyFilter"`
}
