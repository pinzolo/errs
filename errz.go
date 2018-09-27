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
func New(message string) error {
	return &wrapper{
		err: errors.New(message),
		pcs: pcs(defaultSkip),
	}
}

// Errorf returns new error annotated message that is build given format and args.
// Returned error has stack trace.
func Errorf(format string, a ...interface{}) error {
	return &wrapper{
		err: fmt.Errorf(format, a...),
		pcs: pcs(defaultSkip),
	}
}

func pcs(skip int) []uintptr {
	v := make([]uintptr, MaxDepth+1)
	n := runtime.Callers(skip, v)
	return v[0:n]
}
