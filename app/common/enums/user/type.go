package userenum

type UserType string

var (
	Internal UserType = "INTERNAL"
	External UserType = "EXTERNAL"
)

var enums = map[UserType]struct{}{
	Internal: {},
	External: {},
}

func IsValidType(s string) bool {
	_, exists := enums[UserType(s)]
	return exists
}
