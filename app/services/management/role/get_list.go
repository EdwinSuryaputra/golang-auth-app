package role

import (
	"context"

	roleInterface "golang-auth-app/app/interfaces/management/role"
)

func (i *impl) GetList(ctx context.Context, payload *roleInterface.AdapterSqlGetPaginatedPayload) (*roleInterface.ServiceGetListRoleResult, error) {
	roles, err := i.roleSqlAdapter.GetPaginated(ctx, payload)
	if err != nil {
		return nil, err
	}

	return &roleInterface.ServiceGetListRoleResult{
		Entries:  roles.Entries,
		TotalRow: roles.TotalRow,
	}, nil
}
