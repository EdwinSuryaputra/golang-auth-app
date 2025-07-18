package httpadapters

import (
	"github.com/go-resty/resty/v2"
	"go.uber.org/fx"
)

var Module = fx.Module("adapters/http",
	fx.Provide(
		initResty,
	),
)

func initResty() *resty.Client {
	return resty.New()
}