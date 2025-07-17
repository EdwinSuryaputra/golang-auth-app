package supplier

import (
	supplierInterface "golang-auth-app/app/interfaces/supplier"

	"github.com/go-resty/resty/v2"
)

type impl struct {
	restyClient *resty.Client
}

func New(
	restyClient *resty.Client,
) supplierInterface.AdapterHttp {
	return &impl{
		restyClient: restyClient,
	}
}
