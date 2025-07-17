package casbin

import (
	"github.com/rotisserie/eris"
)

func (i *impl) LoadPoliciesFromDB() error {
	if err := i.casbinEnforcer.LoadPolicy(); err != nil {
		return eris.Wrap(err, "error occurred during load casbin policies into memory")
	}

	return nil
}
