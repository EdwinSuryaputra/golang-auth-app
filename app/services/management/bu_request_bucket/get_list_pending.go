package burequestbucket

import (
	"context"

	statusenum "golang-auth-app/app/common/enums/status"
	"golang-auth-app/app/datasources/sql/gorm/model"
	burbDto "golang-auth-app/app/interfaces/management/bu_request_bucket/dto"
	userDto "golang-auth-app/app/interfaces/management/user/dto"
	publicfacingutil "golang-auth-app/app/utils/publicfacing"
	sliceutil "golang-auth-app/app/utils/slice"
)

func (i *impl) GetListPending(
	ctx context.Context,
	payload *burbDto.ServiceGetListPendingPayload,
) (*burbDto.ServiceGetListPendingResult, error) {
	var err error

	var users []*model.User
	var userIds []int32
	if payload.Username != "" || payload.Fullname != "" {
		fetchedUsers, err := i.userSqlAdapter.GetPaginated(ctx, &userDto.AdapterSqlGetPaginatedPayload{
			Username: payload.Username,
			Fullname: payload.Fullname,
			Page:     1,
			Limit:    30,
		})
		if err != nil {
			return nil, err
		} else if len(fetchedUsers.Entries) < 1 {
			return &burbDto.ServiceGetListPendingResult{}, nil
		}

		users = sliceutil.Map(fetchedUsers.Entries, func(dt *userDto.AdapterSqlGetPaginatedEntry) *model.User {
			userIds = append(userIds, dt.ID)

			return &model.User{
				ID:       dt.ID,
				Username: dt.Username,
				FullName: dt.FullName,
			}
		})
	}

	burbPaginated, err := i.burbSqlAdapter.GetPaginated(ctx, &burbDto.AdapterSqlGetPaginatedPayload{
		RequestDate: payload.RequestDate,
		UserIds:     userIds,
		BuLevel:     payload.BuLevel,
		BuLocation:  payload.BuLocation,
		Status:      string(statusenum.PendingApproval),
		Page:        payload.Page,
		Limit:       payload.Limit,
	})
	if err != nil {
		return nil, err
	}

	if len(users) < 1 {
		users, err = i.userSqlAdapter.GetUsersByIds(ctx, sliceutil.Map(burbPaginated.Entries, func(dt *burbDto.AdapterSqlGetPaginatedEntry) int32 {
			return dt.UserId
		}))
		if err != nil {
			return nil, err
		}
	}
	usersPerId := sliceutil.AssociateBy(users, func(dt *model.User) int32 { return dt.ID })

	entries := []*burbDto.ServiceGetListPendingEntry{}
	for _, bu := range burbPaginated.Entries {
		encodedId, err := publicfacingutil.Encode(bu.Id)
		if err != nil {
			return nil, err
		}

		entry := &burbDto.ServiceGetListPendingEntry{
			Id:          encodedId,
			RequestDate: bu.RequestDate,
			BuLevel:     bu.BuLevel,
			BuLocation:  bu.BuLocation,
		}

		if user, isExist := usersPerId[bu.UserId]; isExist {
			entry.UserId, err = publicfacingutil.Encode(bu.UserId)
			if err != nil {
				return nil, err
			}

			entry.Username = user.Username
			entry.Fullname = user.FullName
		}

		entries = append(entries, entry)
	}

	return &burbDto.ServiceGetListPendingResult{
		Entries:  entries,
		TotalRow: burbPaginated.TotalRow,
	}, nil
}
