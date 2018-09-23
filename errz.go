package errz

import (
	"errors"
	"fmt"
	"runtime"
)

// MaxDepth is max depth of stack trace.
var MaxDepth = 32

const defaultSkip = 3

// New error that annotated given message.
// Returned error has stack trace.
func New(msg string) error {
	err := errors.New(msg)
	return &box{
		msg:   msg,
		trace: newTrace(defaultSkip),
		cause: newCause(err, defaultSkip),
	}
}

// Errorf returns new error annotated message that is build given format and args.
// Returned error has stack trace.
func Errorf(format string, a ...interface{}) error {
	err := fmt.Errorf(format, a...)
	return &box{
		msg:   err.Error(),
		trace: newTrace(defaultSkip),
		cause: newCause(err, defaultSkip),
	}
}

func pcs(skip int) []uintptr {
	v := make([]uintptr, MaxDepth+1)
	n := runtime.Callers(skip, v)
	return v[0:n]
}
