package roleenum

type RoleType string

var (
	Internal RoleType = "INTERNAL"
	External RoleType = "EXTERNAL"
)

var enums = map[RoleType]struct{}{
	Internal: {},
	External: {},
}

func IsValidRoleType(s string) bool {
	_, exists := enums[RoleType(s)]
	return exists
}
