package errorcode

// errorItf is an interface that represents an error code.
// Custom error codes format for messages and code representation should implement this interface.
type errorItf interface {
	GetHttpStatusCode() int
	GetErrEnum() string
	Error() string
}

// Error is the standard error implementation of the ErrorItf interface.
type Error struct {
	ErrHttpStatusCode int
	ErrEnum           string
	ErrMessage        string
}

func (e Error) GetHttpStatusCode() int {
	return e.ErrHttpStatusCode
}

func (e Error) GetErrEnum() string {
	return e.ErrEnum
}

func (e Error) Error() string {
	return e.ErrMessage
}
