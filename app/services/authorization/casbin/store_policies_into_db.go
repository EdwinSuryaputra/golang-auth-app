package casbin

import (
	casbinDto "golang-auth-app/app/interfaces/authorization/casbin/dto"

	"github.com/rotisserie/eris"
)

func (i *impl) StorePoliciesIntoDB(
	payloads []*casbinDto.AuthorizePolicyPayload,
) error {
	for _, payload := range payloads {
		for _, role := range payload.AuthorizedUserRoles {
			_, err := i.casbinEnforcer.AddPolicy(role, payload.RequiredResource)
			if err != nil {
				return eris.Wrap(err, "error occured during casbin storePolicies")
			}
		}
	}

	return nil
}
