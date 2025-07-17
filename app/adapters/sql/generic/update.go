package generic

import (
	"context"
)

func (i *impl) Update(ctx context.Context, data interface{}, conditions map[string]interface{}) error {
	q := i.db.Model(new(interface{}))

	for key, value := range conditions {
		q = q.Where(key, value)
	}

	return q.Updates(&data).Error
}
