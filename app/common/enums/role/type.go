package roleenum

type RoleType string

var (
	A RoleType = "A"
	B RoleType = "B"
)

var enums = map[RoleType]struct{}{
	A: {},
	B: {},
}

func IsValidRoleType(s string) bool {
	_, exists := enums[RoleType(s)]
	return exists
}
