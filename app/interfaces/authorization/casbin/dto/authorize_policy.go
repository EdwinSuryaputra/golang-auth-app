package dto

type AuthorizePolicyPayload struct {
	RequiredResource    string
	AuthorizedUserRoles []string
}
