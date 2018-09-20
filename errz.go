package errz

import (
	"fmt"
)

type base struct {
	msg string
	st  *StackTrace
}

func (b *base) Error() string {
	return b.msg
}

func (b *base) Trace() *StackTrace {
	return b.st
}

// New error that annotated given message.
// Returned error has stack trace.
func New(msg string) error {
	return &base{
		msg: msg,
		st:  callers(),
	}
}

// Errorf returns new error annotated message that is build given format and args.
// Returned error has stack trace.
func Errorf(format string, a ...interface{}) error {
	return &base{
		msg: fmt.Sprintf(format, a...),
		st:  callers(),
	}
}
