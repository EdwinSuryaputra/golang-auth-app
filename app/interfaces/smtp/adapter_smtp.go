package smtp

import "context"

type AdapterSMTP interface {
	SendEmail(ctx context.Context, payload *SendEmailPayload) error
}
