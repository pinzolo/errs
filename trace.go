package errz

import (
	"runtime"
	"strings"
)

// MaxDepth is max depth of stack trace.
var MaxDepth = 32

// drop entrance and constructor funcs.
const firstDepth = 2

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

func callers() *StackTrace {
	i := firstDepth
	st := StackTrace{
		Callers: make([]*Caller, 0),
	}
	for {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if i == MaxDepth+firstDepth {
			st.More = true
			break
		}
		fn := funcName(pc)
		c := &Caller{
			PC:   pc,
			File: file,
			Line: line,
			Func: fn,
		}
		st.Callers = append(st.Callers, c)
		i++
	}
	return &st
}

func funcName(pc uintptr) string {
	name := runtime.FuncForPC(pc).Name()
	i := strings.LastIndex(name, "/")
	name = name[i+1:]
	i = strings.Index(name, ".")
	return name[i+1:]
}
