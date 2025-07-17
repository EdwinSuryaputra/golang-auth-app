package casbin

import (
	"context"
)

func (i *impl) SyncAllPolicies(ctx context.Context) error {
	roleResourceMappings, err := i.roleSqlAdapter.GetRoleResourceMappings(ctx, nil)
	if err != nil {
		return err
	}

	policies, err := i.GetAuthorizePolicyPayload(ctx, roleResourceMappings)
	if err != nil {
		return err
	}

	i.casbinEnforcer.ClearPolicy()

	err = i.StorePoliciesIntoDB(policies)
	if err != nil {
		return err
	}

	err = i.LoadPoliciesFromDB()
	if err != nil {
		return err
	}

	return nil
}
