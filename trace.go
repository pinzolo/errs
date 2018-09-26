package errz

import (
	"fmt"
	"io"
	"runtime"
	"strconv"
	"strings"
)

// Caller is line of stack trace.
type Caller struct {
	PC      uintptr
	File    string
	Line    int
	Package string
	Func    string
}

// FullFunc is package + function name that separated by period.
func (c *Caller) FullFunc() string {
	if c.Package == "" {
		return c.Func
	}
	return c.Package + "." + c.Func
}

// Format the caller according to the fmt.Formatter interface.
//
//   %s    source file
//   %v    equivalent to %s:%d
//   %c    compatible with pkg/errors(%v)
//   %+s   function name and source file separated by '\t'
//   %+v   equivalent to %+s:%d
//   %+c   compatible with pkg/errors(%+v)
// nolint: gas
func (c *Caller) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		if c.Func == "" {
			io.WriteString(s, "<unknown>")
		} else {
			if s.Flag('+') {
				fmt.Fprintf(s, "%s\t%s", c.FullFunc(), c.File)
			} else {
				io.WriteString(s, c.File)
			}
		}
	case 'c':
		if c.Func == "" {
			io.WriteString(s, "unknown")
		} else {
			if s.Flag('+') {
				fmt.Fprintf(s, "%s\n\t%s", c.FullFunc(), c.File)
			} else {
				io.WriteString(s, c.File)
			}
		}
		io.WriteString(s, ":")
		io.WriteString(s, strconv.Itoa(c.Line))
	case 'v':
		c.Format(s, 's')
		io.WriteString(s, ":")
		c.Format(s, 'd')
	}
}

// StackTrace is stack traces and more traces.
type StackTrace struct {
	More    bool
	Callers []*Caller
}

// Format the stack trace according to the fmt.Formatter interface.
// nolint: gas
func (st *StackTrace) Format(s fmt.State, verb rune) {
	format := callerFormat(s, verb)
	switch verb {
	case 's':
		for _, c := range st.Callers {
			fmt.Fprintf(s, format, c)
		}
	case 'v':
		for _, c := range st.Callers {
			fmt.Fprintf(s, format, c)
		}
		if s.Flag('+') && st.More {
			io.WriteString(s, "\n\tand more...")
		}
	case 'c':
		for _, c := range st.Callers {
			fmt.Fprintf(s, format, c)
		}
	}
}

// Trace returns stack trace.
// If err is nil, returns nil.
func Trace(err error) *StackTrace {
	if err == nil {
		return nil
	}

	if w, ok := err.(*wrapper); ok {
		return w.Trace()
	}
	return nil
}

// CauseTrace returns stack trace of original error.
// If given error is not wrapped, CauseTrace returns nil.
// If given error is nil, CauseTrace returns nil.
func CauseTrace(err error) *StackTrace {
	if err == nil {
		return nil
	}

	var st *StackTrace
	if w, ok := err.(*wrapper); ok {
		st = w.Trace()
		if st2 := Trace(w.err); st2 != nil {
			st = st2
		}
	}

	return st
}

func callers(pcs []uintptr) []*Caller {
	cs := make([]*Caller, len(pcs))
	for i, pc := range pcs {
		fn := runtime.FuncForPC(pc)
		var pkg, name string
		if fn != nil {
			pkg, name = separateFuncName(fn.Name())
		}
		file, line := fn.FileLine(pc)
		cs[i] = &Caller{
			PC:      pc,
			File:    file,
			Line:    line,
			Package: pkg,
			Func:    name,
		}
	}
	return cs
}

func separateFuncName(path string) (pkg, name string) {
	i := strings.LastIndex(path, "/")
	if i != -1 {
		pkg = path[:i]
	}
	name = path[i+1:]
	i = strings.Index(name, ".")
	if i != -1 {
		if pkg != "" {
			pkg += "/"
		}
		pkg += name[:i]
	}
	name = name[i+1:]
	return
}

func callerFormat(s fmt.State, verb rune) string {
	prefix := "\n\t"
	if verb == 'c' {
		prefix = "\n"
	}
	if s.Flag('+') {
		return prefix + "%+" + string(verb)
	}
	return prefix + "%" + string(verb)
}
