package applicationenum

type Name string

const (
	A Name = "A"
)

var enums = map[Name]struct{}{
	A: {},
}

func IsValidEnum(s string) bool {
	_, exists := enums[Name(s)]
	return exists
}
