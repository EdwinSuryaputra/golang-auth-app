package ctxutil

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

func SyncLocalsToContext(c *fiber.Ctx, keys ...string) {
	ctx := c.UserContext()
	for _, key := range keys {
		val := c.Locals(key)
		ctx = context.WithValue(ctx, key, val)
	}
	c.SetUserContext(ctx)
}
