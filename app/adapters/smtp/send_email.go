package smtpadapters

import (
	"context"
	"fmt"
	"net/smtp"
	"strings"

	smtpItf "golang-auth-app/app/interfaces/smtp"

	"github.com/rotisserie/eris"
)

func (i *impl) SendEmail(ctx context.Context, payload *smtpItf.SendEmailPayload) error {
	msg := fmt.Sprintf("To: %s\r\nSubject: %s\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=\"UTF-8\"\r\n\r\n%s",
		strings.Join(payload.Recipients, ", "),
		payload.Subject,
		payload.HTMLBody,
	)

	auth := smtp.PlainAuth("", payload.User, payload.Password, payload.Host)

	err := smtp.SendMail(fmt.Sprintf("%s:%s", payload.Host, payload.Port), auth, payload.User, payload.Recipients, []byte(msg))
	if err != nil {
		return eris.Wrap(err, err.Error())
	}

	return nil
}
