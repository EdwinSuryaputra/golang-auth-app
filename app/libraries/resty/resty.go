package resty

import (
	"go.uber.org/fx"

	"github.com/go-resty/resty/v2"
)

var Module = fx.Module("libraries/resty", fx.Provide(
	initResty,
))

func initResty() *resty.Client {
	return resty.New()
}
