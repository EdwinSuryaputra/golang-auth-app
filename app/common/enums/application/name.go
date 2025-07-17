package applicationenum

type Name string

const (
	Nte Name = "NTE"
	PRN Name = "PRN"
	HRN Name = "HRN"
)

var enums = map[Name]struct{}{
	Nte: {},
	PRN: {},
	HRN: {},
}

func IsValidEnum(s string) bool {
	_, exists := enums[Name(s)]
	return exists
}
