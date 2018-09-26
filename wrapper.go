package errz

import (
	"fmt"
	"io"
)

type wrapper struct {
	err  error
	pcs  []uintptr
	st   *StackTrace
	code string
	msg  string
}

func (w *wrapper) Error() string {
	msg := w.err.Error()
	if w.msg != "" {
		msg = w.msg + ": " + msg
	}
	if w.code != "" {
		msg = "[" + w.code + "] " + msg
	}
	return msg
}

func (w *wrapper) Trace() *StackTrace {
	if w.st != nil {
		return w.st
	}

	w.st = &StackTrace{}
	cs := callers(w.pcs)
	if len(cs) > MaxDepth {
		w.st.More = true
		cs = cs[0:MaxDepth]
	}
	w.st.Callers = cs
	return w.st
}

func (w *wrapper) message() string {
	if w.code != "" {
		return "[" + w.code + "] " + w.msg
	}
	return w.msg
}

func (w *wrapper) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		if s.Flag('+') {
			w.formatPlusS(s)
		} else {
			io.WriteString(s, w.Error())
		}
	case 'v':
		if s.Flag('+') {
			w.formatPlusV(s)
		} else {
			io.WriteString(s, w.Error())
		}
	case 'c':
		if s.Flag('+') {
			w.formatPlusC(s)
		} else {
			io.WriteString(s, w.Error())
		}
	}
}

func (w *wrapper) formatPlusS(s io.Writer) {
	io.WriteString(s, w.Error())
	fmt.Fprintf(s, "%+s", w.Trace())
	if w.err != nil {
		if iw, ok := w.err.(*wrapper); ok {
			fmt.Fprintf(s, "%+s", iw)
		}
	}
}

func (w *wrapper) formatPlusV(s io.Writer) {
	io.WriteString(s, w.Error())
	fmt.Fprintf(s, "%+v", w.Trace())
	if w.err != nil {
		if iw, ok := w.err.(*wrapper); ok {
			fmt.Fprintf(s, "%+v", iw)
		}
	}
}

func (w *wrapper) formatPlusC(s io.Writer) {
	if w.err != nil {
		if iw, ok := w.err.(*wrapper); ok {
			fmt.Fprintf(s, "%+c", iw)
		} else {
			fmt.Fprintf(s, "%+v", w.err)
		}
	}
	io.WriteString(s, w.message())
	fmt.Fprintf(s, "%+c", w.Trace())
}
