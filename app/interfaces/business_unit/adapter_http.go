package businessunit

import (
	"context"

	businessunitmodel "golang-auth-app/app/common/models/business_unit"
)

type AdapterHttp interface {
	GetBusinessUnits(ctx context.Context, businessUnitLevel string, businessUnitId string) ([]*businessunitmodel.BusinessUnit, error)
}
