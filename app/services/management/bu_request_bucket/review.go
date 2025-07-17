package burequestbucket

import (
	"context"
	"fmt"
	"time"

	reviewenum "golang-auth-app/app/common/enums/review"
	statusenum "golang-auth-app/app/common/enums/status"
	"golang-auth-app/app/datasources/sql/gorm/model"
	"golang-auth-app/app/interfaces/errorcode"
	burbDto "golang-auth-app/app/interfaces/management/bu_request_bucket/dto"
)

func (i *impl) Review(ctx context.Context, payload *burbDto.ServiceReviewPayload) error {
	now := time.Now()

	existingBurb, err := i.burbSqlAdapter.GetBuRequestBucketById(ctx, payload.Id)
	if err != nil {
		return err
	}

	newDataBurb := *existingBurb
	newDataBurb.ReviewedAt = &now
	newDataBurb.ReviewedBy = &payload.Modifier
	newDataBurb.UpdatedAt = now
	newDataBurb.UpdatedBy = payload.Modifier

	existingUser, err := i.userSqlAdapter.GetUserById(ctx, existingBurb.UserID)
	if err != nil {
		return err
	}

	toBeUpdatedUser := &model.User{
		ID:        existingBurb.UserID,
		UpdatedBy: payload.Modifier,
		UpdatedAt: now,
	}

	/*
		Flows
		1. Action = Approve, currentStatus = PENDING_APPROVAL, user bu status => ACTIVE
		2. Action = Reject, currentStatus = PENDING_APPROVAL, user bu status => UNASSIGNED
		3. Action = Reject, currentStatus = ACTIVE_PENDING_APPROVAL, user bu status => ACTIVE
	*/

	currentUserBurbStatus := statusenum.Status(*existingUser.BusinessUnitAssignmentStatus)
	var newBurbStatus statusenum.Status
	var activityLogMessage string

	switch reviewenum.Action(payload.Action) {
	case reviewenum.Approve:
		newBurbStatus = statusenum.Approved

		newUserBurbStatus := statusenum.Active.ToString()
		toBeUpdatedUser.BusinessUnitAssignmentStatus = &newUserBurbStatus

		switch currentUserBurbStatus {
		case statusenum.PendingApproval:
			activityLogMessage = fmt.Sprintf("%s approved %s BU assignment", payload.Modifier, existingUser.Username)
		case statusenum.ActivePendingApproval:
			activityLogMessage = fmt.Sprintf("%s approved %s BU assignment changes", payload.Modifier, existingUser.Username)
		}
	case reviewenum.Reject:
		newBurbStatus = statusenum.Rejected

		switch currentUserBurbStatus {
		case statusenum.PendingApproval:
			activityLogMessage = fmt.Sprintf("%s rejected %s BU assignment", payload.Modifier, existingUser.Username)
		case statusenum.ActivePendingApproval:
			activityLogMessage = fmt.Sprintf("%s rejected %s BU assignment changes", payload.Modifier, existingUser.Username)
		}
	}
	newDataBurb.Status = newBurbStatus.ToString()

	switch reviewenum.Action(payload.Action) {
	case reviewenum.Reject:
		existingTempUser, err := i.userSqlAdapter.GetTempUserByUserId(ctx, existingUser.ID)
		if err != nil {
			return err
		} else if existingTempUser == nil {
			return errorcode.WithCustomMessage(errorcode.ErrCodeNotFound, "temp user is not found")
		}

		toBeUpdatedUser.BusinessUnitLevel = existingTempUser.BusinessUnitLevel
		toBeUpdatedUser.BusinessUnitLocationID = existingTempUser.BusinessUnitLocationID
		toBeUpdatedUser.BusinessUnitLocation = existingTempUser.BusinessUnitLocation
		toBeUpdatedUser.BusinessUnitAssignmentStatus = existingTempUser.BusinessUnitAssignmentStatus
	}

	if err = i.burbSqlAdapter.UpdateBuRequestBucket(ctx, &newDataBurb); err != nil {
		return err
	}

	if err = i.userSqlAdapter.UpdateUser(ctx, toBeUpdatedUser); err != nil {
		return err
	}

	if activityLogMessage != "" && existingUser.ActivityLogID != nil {
		err = i.activityLogHttpAdapter.Insert(ctx, *existingUser.ActivityLogID, activityLogMessage, "SUCCESS")
		if err != nil {
			return err
		}
	}

	return nil
}
