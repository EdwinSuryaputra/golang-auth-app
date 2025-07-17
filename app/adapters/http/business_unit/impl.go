package businessunit

import (
	businessunit "golang-auth-app/app/interfaces/business_unit"

	"github.com/go-resty/resty/v2"
)

type impl struct {
	restyClient *resty.Client
}

func New(
	restyClient *resty.Client,
) businessunit.AdapterHttp {
	return &impl{
		restyClient: restyClient,
	}
}
