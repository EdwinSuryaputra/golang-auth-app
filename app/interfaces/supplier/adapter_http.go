package supplier

import (
	"context"

	suppliermodel "golang-auth-app/app/common/models/supplier"
)

type AdapterHttp interface {
	GetSupplier(ctx context.Context, supplierId string, supplierName string) ([]*suppliermodel.Supplier, error)
}
