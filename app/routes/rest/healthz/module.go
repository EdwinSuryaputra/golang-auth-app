package healthz

import "go.uber.org/fx"

var prefix = "/healthz"

var Module = fx.Module("routes/rest/healthz", fx.Invoke(
	getHealthz,
))
