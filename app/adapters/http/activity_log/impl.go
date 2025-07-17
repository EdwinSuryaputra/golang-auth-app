package activitylog

import (
	activitylog "golang-auth-app/app/interfaces/activity_log"

	"github.com/go-resty/resty/v2"
)

type impl struct {
	restyClient *resty.Client
}

func New(
	restyClient *resty.Client,
) activitylog.AdapterHttp {
	return &impl{
		restyClient: restyClient,
	}
}
