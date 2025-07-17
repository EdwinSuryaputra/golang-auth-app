package generic

import (
	"context"
)

func (i *impl) Insert(ctx context.Context, data interface{}) error {
	return i.db.Create(data).Error
}