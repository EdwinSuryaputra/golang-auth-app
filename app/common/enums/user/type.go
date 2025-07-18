package userenum

type UserType string

var (
	A UserType = "A"
	B UserType = "B"
)

var enums = map[UserType]struct{}{
	A: {},
	B: {},
}

func IsValidType(s string) bool {
	_, exists := enums[UserType(s)]
	return exists
}
