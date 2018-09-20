package errz_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/pinzolo/errz"
)

func TestTrace(t *testing.T) {
	fns := []struct {
		fn   func() error
		name string
		line int
		note string
	}{
		{
			func() error {
				return errz.New("error")
			},
			"func1",
			20,
			"errz.New",
		},
		{
			func() error {
				return errz.Errorf("error %s", "format")
			},
			"func2",
			28,
			"errz.Errorf",
		},
		{
			func() error {
				return errz.Wrap(errors.New("original"), "wrap")
			},
			"func3",
			36,
			"errz.Wrap",
		},
		{
			func() error {
				return errz.Wrapf(errors.New("original"), "wrap %s", "format")
			},
			"func4",
			44,
			"errz.Wrapf",
		},
	}

	for _, fn := range fns {
		t.Run(fn.note, func(t *testing.T) {
			err := f(4, fn.fn)
			st := errz.Trace(err)
			data := []struct {
				fn   string
				line int
			}{
				{"TestTrace." + fn.name, fn.line},
				{"f4", 36},
				{"f3", 31},
				{"f2", 24},
				{"f1", 17},
				{"f", 10},
			}
			for i, d := range data {
				c := st.Callers[i]
				if !strings.HasSuffix(c.Func, d.fn) {
					t.Errorf("func: want %s, got %s", d.fn, c.Func)
				}
				if d.line != c.Line {
					t.Errorf("line: want %d, got %d", d.line, c.Line)
				}
			}
			if st.More {
				t.Error("real stack trace length is more than max depth, st.More should be false")
			}
		})
	}
}

func TestTraceWithNil(t *testing.T) {
	if errz.Trace(nil) != nil {
		t.Errorf("Trace should return nil when given nil")
	}
}

func TestTraceWithRawError(t *testing.T) {
	if errz.Trace(errors.New("raw error")) != nil {
		t.Errorf("Trace should return nil when given raw error")
	}
}

func TestMoreTrace(t *testing.T) {
	data := []struct {
		depth int
		more  bool
		note  string
	}{
		{8, true, "less"},
		{9, false, "just"},
		{10, false, "more"},
	}
	for _, d := range data {
		t.Run(d.note, func(t *testing.T) {
			defer func() func() {
				max := errz.MaxDepth
				errz.MaxDepth = d.depth
				return func() {
					errz.MaxDepth = max
				}
			}()()
			err := f(4, func() error {
				return errz.New("error")
			})
			st := errz.Trace(err)
			if st.More != d.more {
				t.Errorf("more: want %v, got %v", d.more, st.More)
			}
		})
	}
}
