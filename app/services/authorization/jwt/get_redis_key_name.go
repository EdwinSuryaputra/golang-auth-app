package jwt

import "fmt"

func (i *impl) getRedisKeyName(username string) string {
	return fmt.Sprintf("bearer-token-%s", username)
}
