package user

import "go.uber.org/fx"

var prefix = "/management/user"

var Module = fx.Module("routes/rest/management/user", fx.Invoke(
	getList,
	getDetail,
	createDraft,
	update,
	delete,
	review,
))
