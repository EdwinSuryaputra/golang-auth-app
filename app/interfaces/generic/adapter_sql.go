package generic

import "context"

type AdapterSQL interface {
	Insert(ctx context.Context, data interface{}) error
	Update(ctx context.Context, data interface{}, conditions map[string]interface{}) error
	SoftDelete(ctx context.Context, conditions map[string]interface{}, modifier string) error
}
