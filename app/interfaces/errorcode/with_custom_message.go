package errorcode

import (
	"github.com/rotisserie/eris"
)

func WithCustomMessage(err errorItf, msg string) error {
	return eris.Wrap(
		Error{
			ErrHttpStatusCode: err.GetHttpStatusCode(),
			ErrEnum:           err.GetErrEnum(),
			ErrMessage:        msg,
		}, msg,
	)
}
