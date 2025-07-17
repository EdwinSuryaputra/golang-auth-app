package user

import (
	"context"

	userDto "golang-auth-app/app/interfaces/management/user/dto"
)

func (i *impl) GetList(ctx context.Context, payload *userDto.ServiceGetListPayload) (*userDto.ServiceGetListResult, error) {
	paginatedUsers, err := i.userSqlAdapter.GetPaginated(ctx, payload.Query)
	if err != nil {
		return nil, err
	}

	return &userDto.ServiceGetListResult{
		Entries:  paginatedUsers.Entries,
		TotalRow: paginatedUsers.TotalRow,
	}, nil
}
