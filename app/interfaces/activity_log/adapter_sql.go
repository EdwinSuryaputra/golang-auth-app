package activitylog

import "context"

type AdapterSql interface {
	Insert(ctx context.Context, activityLogId, message, status string) error
}
