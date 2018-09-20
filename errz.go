package errz

import (
	"fmt"
	"runtime"
)

// MaxDepth is max depth of stack trace.
var MaxDepth = 32

type base struct {
	msg string
	pcs []uintptr
	st  *StackTrace
}

func (b *base) Error() string {
	return b.msg
}

func (b *base) Trace() *StackTrace {
	if b.st != nil {
		return b.st
	}

	b.st = &StackTrace{}
	cs := callers(b.pcs)
	if len(cs) > MaxDepth {
		b.st.More = true
		cs = cs[0:MaxDepth]
	}
	b.st.Callers = cs
	return b.st
}

// New error that annotated given message.
// Returned error has stack trace.
func New(msg string) error {
	return &base{
		msg: msg,
		pcs: pcs(),
	}
}

// Errorf returns new error annotated message that is build given format and args.
// Returned error has stack trace.
func Errorf(format string, a ...interface{}) error {
	return &base{
		msg: fmt.Sprintf(format, a...),
		pcs: pcs(),
	}
}

func pcs() []uintptr {
	v := make([]uintptr, MaxDepth+1)
	n := runtime.Callers(3, v)
	return v[0:n]
}
