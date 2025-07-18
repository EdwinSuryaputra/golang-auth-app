package generic

import "context"

type AdapterSQL interface {
	SoftDelete(ctx context.Context, conditions map[string]interface{}, modifier string) error
}
