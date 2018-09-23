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
		memo string
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
		{
			func() error {
				err := errz.New("original")
				err = errz.Wrap(err, "wrap1")
				err = errz.WithCode(err, "E01")
				return errz.Wrapf(err, "wrap%d", 2)
			},
			"func5",
			55,
			"New -> Wrap -> WithCode -> Wrapf",
		},
	}

	for _, fn := range fns {
		t.Run(fn.memo, func(t *testing.T) {
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

func TestTrace_WithNil(t *testing.T) {
	if errz.Trace(nil) != nil {
		t.Errorf("Trace should return nil when given nil")
	}
}

func TestTrace_WithRawError(t *testing.T) {
	if errz.Trace(errors.New("raw error")) != nil {
		t.Errorf("Trace should return nil when given raw error")
	}
}

func TestMoreTrace(t *testing.T) {
	data := []struct {
		depth int
		more  bool
		memo  string
	}{
		{8, true, "less"},
		{9, false, "just"},
		{10, false, "more"},
	}
	for _, d := range data {
		t.Run(d.memo, func(t *testing.T) {
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

func TestCauseTrace(t *testing.T) {
	err := f(4, func() error {
		return errz.New("original")
	})
	err = f(2, func() error {
		return errz.Wrap(err, "wrap")
	})

	data := []struct {
		file string
		line int
		fn   string
	}{
		{"trace_test.go", 141, "TestCauseTrace.func2"},
		{"errz_test.go", 22, "f2"},
		{"errz_test.go", 17, "f1"},
		{"errz_test.go", 10, "f"},
	}
	st := errz.Trace(err)
	for i, d := range data {
		c := st.Callers[i]
		if !strings.HasSuffix(c.File, d.file) {
			t.Errorf("file: want %s, got %s", d.file, c.File)
		}
		if d.line != c.Line {
			t.Errorf("line: want %d, got %d", d.line, c.Line)
		}
		if c.Func != d.fn {
			t.Errorf("func: want %s, got %s", d.fn, c.Func)
		}
	}

	data = []struct {
		file string
		line int
		fn   string
	}{
		{"trace_test.go", 138, "TestCauseTrace.func1"},
		{"errz_test.go", 36, "f4"},
		{"errz_test.go", 31, "f3"},
		{"errz_test.go", 24, "f2"},
		{"errz_test.go", 17, "f1"},
		{"errz_test.go", 10, "f"},
	}
	st = errz.CauseTrace(err)
	for i, d := range data {
		c := st.Callers[i]
		if !strings.HasSuffix(c.File, d.file) {
			t.Errorf("file: want %s, got %s", d.file, c.File)
		}
		if d.line != c.Line {
			t.Errorf("line: want %d, got %d", d.line, c.Line)
		}
		if c.Func != d.fn {
			t.Errorf("func: want %s, got %s", d.fn, c.Func)
		}
	}
}

func TestCauseTrace_NilPattern(t *testing.T) {
	data := []struct {
		err  error
		memo string
	}{
		{nil, "nil"},
		{errors.New("raw"), "raw error"},
	}
	for _, d := range data {
		t.Run(d.memo, func(t *testing.T) {
			if errz.CauseTrace(d.err) != nil {
				t.Errorf("CauseTrace should return nil when given error is %s", d.memo)
			}
		})
	}
}

func TestHasTrace(t *testing.T) {
	orig := errors.New("original")
	data := []struct {
		err  error
		want bool
		memo string
	}{
		{nil, false, "nil"},
		{orig, false, "raw error"},
		{errz.Wrap(orig, "wrap"), true, "wrap"},
		{errz.Wrap(errz.Wrap(errz.Wrap(orig, "wrap1"), "wrap2"), "wrap3"), true, "nested wrap"},
	}

	for _, d := range data {
		t.Run(d.memo, func(t *testing.T) {
			if errz.HasTrace(d.err) != d.want {
				if d.want {
					t.Error("HasTrace should return true if error is wrapped")
				} else {
					t.Error("HasTrace should return false if error is not wrapped")
				}
			}
		})
	}
}
