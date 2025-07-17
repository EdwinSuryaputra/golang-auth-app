package burequestbucket

import "go.uber.org/fx"

var prefix = "/management/bu_request_bucket"

var Module = fx.Module("routes/rest/management/bu_request_bucket", fx.Invoke(
	getListPending,
	getListCompleted,
	review,
))
