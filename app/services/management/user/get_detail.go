package user

import (
	"context"
	"encoding/json"
	"strings"

	userenum "golang-auth-app/app/common/enums/user"

	"golang-auth-app/app/datasources/sql/gorm/model"

	userDto "golang-auth-app/app/interfaces/management/user/dto"

	objectutil "golang-auth-app/app/utils/object"
	publicfacingutil "golang-auth-app/app/utils/publicfacing"
	sliceutil "golang-auth-app/app/utils/slice"

	"github.com/rotisserie/eris"
)

func (i *impl) GetDetail(
	ctx context.Context,
	encodedUserId string,
) (*userDto.ServiceGetDetailResult, error) {
	userId, err := publicfacingutil.Decode(strings.TrimSpace(encodedUserId))
	if err != nil {
		return nil, err
	}

	existingUser, err := i.userSqlAdapter.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	var assignedRoles []*userDto.ServiceGetDetailAssignedRole
	if existingUser.AssignedRoles != nil {
		var userRoles []*model.UserRoleMapping
		if err = json.Unmarshal([]byte(*existingUser.AssignedRoles), &userRoles); err != nil {
			return nil, eris.Wrap(err, err.Error())
		}

		roles, err := i.roleSqlAdapter.GetRolesByIds(ctx, objectutil.Keys(sliceutil.AssociateBy(userRoles, func(dt *model.UserRoleMapping) int32 { return *dt.RoleID })))
		if err != nil {
			return nil, err
		}

		assignedRoles, err = sliceutil.MapWithError(roles, func(dt *model.Role) (*userDto.ServiceGetDetailAssignedRole, error) {
			encodedId, err := publicfacingutil.Encode(dt.ID)
			if err != nil {
				return nil, err
			}

			return &userDto.ServiceGetDetailAssignedRole{
				Id:   encodedId,
				Name: dt.Name,
			}, nil
		})
		if err != nil {
			return nil, err
		}
	}

	result := &userDto.ServiceGetDetailResult{
		Id:            encodedUserId,
		Username:      existingUser.Username,
		Fullname:      existingUser.FullName,
		Email:         existingUser.Email,
		Type:          existingUser.Type,
		Status:        existingUser.Status,
		StartDate:     existingUser.StartDate,
		EndDate:       existingUser.EndDate,
		AssignedRoles: assignedRoles,
		BusinessUnit:  nil,
		Supplier:      nil,
		CreatedAt:     existingUser.CreatedAt,
		CreatedBy:     existingUser.CreatedBy,
		UpdatedAt:     existingUser.UpdatedAt,
		UpdatedBy:     existingUser.UpdatedBy,
		ActivityLogId: existingUser.ActivityLogID,
	}

	switch userenum.UserType(existingUser.Type) {
	case userenum.Internal:
		if existingUser.BusinessUnitLevel != nil &&
			existingUser.BusinessUnitLocationID != nil &&
			existingUser.BusinessUnitLocation != nil &&
			existingUser.BusinessUnitAssignmentStatus != nil {
			encodedBuId, err := publicfacingutil.Encode(*existingUser.BusinessUnitLocationID)
			if err != nil {
				return nil, err
			}

			result.BusinessUnit = &userDto.ServiceGetDetailBusinessUnit{
				Id:     encodedBuId,
				Level:  *existingUser.BusinessUnitLevel,
				Name:   *existingUser.BusinessUnitLocation,
				Status: *existingUser.BusinessUnitAssignmentStatus,
			}
		}
	case userenum.External:
		if existingUser.SupplierID != nil && existingUser.SupplierName != nil {
			encodedSupplierId, err := publicfacingutil.Encode(*existingUser.SupplierID)
			if err != nil {
				return nil, err
			}

			result.Supplier = &userDto.ServiceGetDetailSupplier{
				Id:   encodedSupplierId,
				Name: *existingUser.SupplierName,
			}
		}
	}

	return result, nil
}
