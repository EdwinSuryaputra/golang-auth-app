package sqlenum

type Order string

var (
	Asc  Order = "ASC"
	Desc Order = "DESC"
)

var enums = map[Order]struct{}{
	Asc:  {},
	Desc: {},
}

func IsValidOrder(s string) bool {
	_, exists := enums[Order(s)]
	return exists
}
