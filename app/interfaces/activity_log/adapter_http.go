package activitylog

import "context"

type AdapterHttp interface {
	Insert(ctx context.Context, activityLogId, message, status string) error
}
