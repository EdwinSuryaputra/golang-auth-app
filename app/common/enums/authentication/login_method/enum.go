package loginmethodenum

type LoginMethod string

var (
	CredentialBased LoginMethod = "CREDENTIAL"
)

var enums = map[LoginMethod]struct{}{
	CredentialBased: {},
}

func IsValidEnum(s string) bool {
	_, exists := enums[LoginMethod(s)]
	return exists
}
