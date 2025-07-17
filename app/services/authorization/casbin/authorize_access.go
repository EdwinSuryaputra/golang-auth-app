package casbin

import (
	"context"

	casbinDto "golang-auth-app/app/interfaces/authorization/casbin/dto"

	"github.com/rotisserie/eris"
)

func (i *impl) AuthorizeAccess(
	ctx context.Context,
	payload *casbinDto.AuthorizePolicyPayload,
) (bool, error) {
	var isAuthorized bool
	var err error

	for _, userRole := range payload.AuthorizedUserRoles {
		isAuthorized, err = i.casbinEnforcer.Enforce(userRole, payload.RequiredResource)
		if err != nil {
			return false, eris.Wrap(err, "error occurred during enforce auth access")
		}
	}

	return isAuthorized, nil
}
