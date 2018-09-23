package errz_test

import (
	"strings"
	"testing"

	"github.com/pinzolo/errz"
)

func TestConst(t *testing.T) {
	const ErrFoo = errz.Const("FOO")
	if s := ErrFoo.Error(); s != "FOO" {
		t.Errorf("invalid error message: want %s, got %s", "FOO", s)
	}

	const ErrBar = errz.Const("BAR %s")
	if s := ErrBar.Errorf("BAZ"); s != "BAR BAZ" {
		t.Errorf("invalid error message: want %s, got %s", "FOO", s)
	}
}

func TestConst_WithTrace(t *testing.T) {
	const ErrFoo = errz.Const("FOO")
	if errz.Trace(ErrFoo) != nil {
		t.Errorf("Const error should not have stack trace")
	}
	err := f(4, func() error {
		return ErrFoo.WithTrace()
	})

	data := []struct {
		file string
		line int
		fn   string
	}{
		{"const_test.go", 28, "TestConst_WithTrace.func1"},
		{"errz_test.go", 36, "f4"},
		{"errz_test.go", 31, "f3"},
		{"errz_test.go", 24, "f2"},
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
}
