package generic

import (
	"context"
	"time"

	"github.com/rotisserie/eris"
)

func (i *impl) SoftDelete(ctx context.Context, conditions map[string]interface{}, modifier string) error {
	q := i.db.Model(new(interface{}))

	for key, value := range conditions {
		q = q.Where(key, value)
	}

	err := q.Updates(map[string]interface{}{
		"deleted_at": time.Now(),
		"deleted_by": modifier,
	}).Error
	if err != nil {
		return eris.Wrap(err, "error occurred during SoftDelete")
	}

	return nil
}
