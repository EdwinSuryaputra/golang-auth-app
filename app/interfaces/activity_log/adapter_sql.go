package activitylog

import "context"

type AdapterSQL interface {
	Insert(ctx context.Context, activityLogId, message, status string) error
}
