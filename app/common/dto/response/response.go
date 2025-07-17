package response

import (
	"fmt"
	"math"
	"time"
)

type GetListAPIResponse[T any] struct {
	DataList    []T              `json:"dataList"`
	Pager       *GetListAPIPager `json:"pager"`
	ProcessTime string           `json:"processTime"`
}

type GetListAPIPager struct {
	TotalData   int  `json:"totalData"`
	TotalPage   int  `json:"totalPage"`
	CurrentPage int  `json:"currentPage"`
	NextPage    int  `json:"nextPage"`
	PrevPage    *int `json:"prevPage"`
}

func SetPager(totalData int64, page int, limit int) *GetListAPIPager {
	totalPage := int(math.Ceil(float64(totalData) / float64(limit)))

	var previousPage *int
	if page-1 > 0 {
		val := page - 1
		previousPage = &val
	}

	return &GetListAPIPager{
		TotalData:   int(totalData),
		TotalPage:   totalPage,
		CurrentPage: page,
		NextPage:    page + 1,
		PrevPage:    previousPage,
	}
}

func GetProcessTime(startTime time.Time) string {
	return fmt.Sprintf("%.3f ms", time.Since(startTime).Seconds()*1000)
}
