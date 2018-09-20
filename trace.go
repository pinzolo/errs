package errz

import (
	"runtime"
	"strings"
)

// Caller is line of stack trace.
type Caller struct {
	PC   uintptr
	File string
	Line int
	Func string
}

// StackTrace is stack traces and more traces.
type StackTrace struct {
	More    bool
	Callers []*Caller
}

// Trace returns stack trace.
// If err is nil, returns nil.
func Trace(err error) *StackTrace {
	if err == nil {
		return nil
	}

	type tracer interface {
		Trace() *StackTrace
	}

	if wt, ok := err.(tracer); ok {
		return wt.Trace()
	}
	return nil
}

func callers(pcs []uintptr) []*Caller {
	cs := make([]*Caller, len(pcs))
	for i, pc := range pcs {
		fn := runtime.FuncForPC(pc)
		var name string
		if fn == nil {
			name = "unknown"
		} else {
			name = extractFuncName(fn.Name())
		}
		file, line := fn.FileLine(pc)
		cs[i] = &Caller{
			PC:   pc,
			File: file,
			Line: line,
			Func: name,
		}
	}
	return cs
}

func extractFuncName(path string) string {
	i := strings.LastIndex(path, "/")
	name := path[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}
