package role

import "go.uber.org/fx"

var prefix = "/management/role"

var Module = fx.Module("routes/rest/management/role", fx.Invoke(
	getList,
	getDetail,
	createDraft,
	update,
	getAllResources,
	deleteDraft,
	review,
))
