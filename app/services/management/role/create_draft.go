package role

import (
	"context"
	"fmt"
	"golang-auth-app/app/adapters/sql/gorm/model"

	"github.com/google/uuid"
)

func (i *impl) CreateDraft(
	ctx context.Context,
	newRole *model.Role,
) error {
	activityLogId := uuid.NewString()
	newRole.ActivityLogID = &activityLogId

	err := i.roleSqlAdapter.InsertRole(ctx, newRole)
	if err != nil {
		return err
	}

	err = i.activityLogSqlAdapter.Insert(ctx, activityLogId, fmt.Sprintf("Role was created draft by %s", newRole.CreatedBy), "SUCCESS")
	if err != nil {
		return err
	}

	return nil
}
