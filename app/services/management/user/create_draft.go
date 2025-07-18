package user

import (
	"context"
	"fmt"
	"golang-auth-app/app/adapters/sql/gorm/model"

	"github.com/google/uuid"
)

func (i *impl) CreateDraft(ctx context.Context, newUser *model.User) error {
	activityLogId := uuid.NewString()
	newUser.ActivityLogID = &activityLogId

	err := i.userSqlAdapter.InsertUser(ctx, newUser)
	if err != nil {
		return err
	}

	err = i.activityLogSqlAdapter.Insert(ctx, activityLogId, fmt.Sprintf("User was created draft by %s", newUser.CreatedBy), "SUCCESS")
	if err != nil {
		return err
	}

	return nil
}
