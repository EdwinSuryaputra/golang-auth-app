package dto

type LoginPayload struct {
	CredentialBasedPayload *CredentialBasedPayload
	IsKeepLoggedIn         bool
}

type CredentialBasedPayload struct {
	Username string
	Password string
}

type LoginResult struct {
	UserId            string
	Username          string
	FullName          string
	TokenString       string
	IsDefaultPassword bool
}
