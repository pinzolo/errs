package errz_test

import (
	"bytes"
	"errors"
	"fmt"
	"testing"

	"github.com/pinzolo/errz"
)

func TestFormatS(t *testing.T) {
	err1 := errz.NewWrapper(errors.New("error"), "", "", newStack(callers1(), false))
	err2 := errz.NewWrapper(err1, "", "wrap1", newStack(callers2(), false))
	err3 := errz.NewWrapper(err2, "", "wrap2", newStack(callers3(), false))

	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%s", err3)
	want := "wrap2: wrap1: error"

	if buf.String() != want {
		t.Error("invalid format result")
		t.Errorf("got: %s", buf.String())
		t.Errorf("want: %s", want)
	}
}

func TestFormatPlusS(t *testing.T) {
	err1 := errz.NewWrapper(errors.New("error"), "", "", newStack(callers1(), false))
	err2 := errz.NewWrapper(err1, "", "wrap1", newStack(callers2(), false))
	err3 := errz.NewWrapper(err2, "", "wrap2", newStack(callers3(), false))

	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%+s", err3)
	want := `wrap2: wrap1: error
	github.com/pinzolo/errz_test.TestFormat.func2	/src/github.com/pinzolo/errz/errz_test.go
	github.com/pinzolo/errz_test.f2	/src/github.com/pinzolo/errz/errz_test.go
	github.com/pinzolo/errz_test.f1	/src/github.com/pinzolo/errz/errz_test.go
	github.com/pinzolo/errz_test.f	/src/github.com/pinzolo/errz/errz_test.go
wrap1: error
	github.com/pinzolo/errz_test.TestFormat.func2.1	/src/github.com/pinzolo/errz/errz_test.go
	github.com/pinzolo/errz_test.f3	/src/github.com/pinzolo/errz/errz_test.go
	github.com/pinzolo/errz_test.f2	/src/github.com/pinzolo/errz/errz_test.go
	github.com/pinzolo/errz_test.f1	/src/github.com/pinzolo/errz/errz_test.go
	github.com/pinzolo/errz_test.f	/src/github.com/pinzolo/errz/errz_test.go
error
	github.com/pinzolo/errz_test.TestFormat.func2.1.1	/src/github.com/pinzolo/errz/errz_test.go
	github.com/pinzolo/errz_test.f4	/src/github.com/pinzolo/errz/errz_test.go
	github.com/pinzolo/errz_test.f3	/src/github.com/pinzolo/errz/errz_test.go
	github.com/pinzolo/errz_test.f2	/src/github.com/pinzolo/errz/errz_test.go
	github.com/pinzolo/errz_test.f1	/src/github.com/pinzolo/errz/errz_test.go
	github.com/pinzolo/errz_test.f	/src/github.com/pinzolo/errz/errz_test.go
`
	if buf.String() != want {
		t.Error("invalid format result")
		t.Errorf("got: %s", buf.String())
		t.Errorf("want: %s", want)
	}
}

func TestFormatV(t *testing.T) {
	err1 := errz.NewWrapper(errors.New("error"), "", "", newStack(callers1(), false))
	err2 := errz.NewWrapper(err1, "", "wrap1", newStack(callers2(), false))
	err3 := errz.NewWrapper(err2, "", "wrap2", newStack(callers3(), false))

	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%v", err3)
	want := "wrap2: wrap1: error"

	if buf.String() != want {
		t.Error("invalid format result")
		t.Errorf("got: %s", buf.String())
		t.Errorf("want: %s", want)
	}
}

func TestFormatPlusV(t *testing.T) {
	err1 := errz.NewWrapper(errors.New("error"), "", "", newStack(callers1(), false))
	err2 := errz.NewWrapper(err1, "", "wrap1", newStack(callers2(), false))
	err3 := errz.NewWrapper(err2, "", "wrap2", newStack(callers3(), false))

	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%+v", err3)
	want := `wrap2: wrap1: error
	github.com/pinzolo/errz_test.TestFormat.func2	/src/github.com/pinzolo/errz/errz_test.go:95
	github.com/pinzolo/errz_test.f2	/src/github.com/pinzolo/errz/errz_test.go:26
	github.com/pinzolo/errz_test.f1	/src/github.com/pinzolo/errz/errz_test.go:19
	github.com/pinzolo/errz_test.f	/src/github.com/pinzolo/errz/errz_test.go:12
wrap1: error
	github.com/pinzolo/errz_test.TestFormat.func2.1	/src/github.com/pinzolo/errz/errz_test.go:96
	github.com/pinzolo/errz_test.f3	/src/github.com/pinzolo/errz/errz_test.go:31
	github.com/pinzolo/errz_test.f2	/src/github.com/pinzolo/errz/errz_test.go:26
	github.com/pinzolo/errz_test.f1	/src/github.com/pinzolo/errz/errz_test.go:19
	github.com/pinzolo/errz_test.f	/src/github.com/pinzolo/errz/errz_test.go:12
error
	github.com/pinzolo/errz_test.TestFormat.func2.1.1	/src/github.com/pinzolo/errz/errz_test.go:97
	github.com/pinzolo/errz_test.f4	/src/github.com/pinzolo/errz/errz_test.go:38
	github.com/pinzolo/errz_test.f3	/src/github.com/pinzolo/errz/errz_test.go:33
	github.com/pinzolo/errz_test.f2	/src/github.com/pinzolo/errz/errz_test.go:24
	github.com/pinzolo/errz_test.f1	/src/github.com/pinzolo/errz/errz_test.go:19
	github.com/pinzolo/errz_test.f	/src/github.com/pinzolo/errz/errz_test.go:12
`
	if buf.String() != want {
		t.Error("invalid format result")
		t.Errorf("got: %s", buf.String())
		t.Errorf("want: %s", want)
	}
}

func TestFormatC(t *testing.T) {
	err1 := errz.NewWrapper(errors.New("error"), "", "", newStack(callers1(), false))
	err2 := errz.NewWrapper(err1, "", "wrap1", newStack(callers2(), false))
	err3 := errz.NewWrapper(err2, "", "wrap2", newStack(callers3(), false))

	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%c", err3)
	want := "wrap2: wrap1: error"

	if buf.String() != want {
		t.Error("invalid format result")
		t.Errorf("got: %s", buf.String())
		t.Errorf("want: %s", want)
	}
}

func TestFormatPlusC(t *testing.T) {
	err1 := errz.NewWrapper(errors.New("error"), "", "", newStack(callers1(), false))
	err2 := errz.NewWrapper(err1, "", "wrap1", newStack(callers2(), false))
	err3 := errz.NewWrapper(err2, "", "wrap2", newStack(callers3(), false))

	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "%+c", err3)
	want := `error
github.com/pinzolo/errz_test.TestFormat.func2.1.1
	/src/github.com/pinzolo/errz/errz_test.go:97
github.com/pinzolo/errz_test.f4
	/src/github.com/pinzolo/errz/errz_test.go:38
github.com/pinzolo/errz_test.f3
	/src/github.com/pinzolo/errz/errz_test.go:33
github.com/pinzolo/errz_test.f2
	/src/github.com/pinzolo/errz/errz_test.go:24
github.com/pinzolo/errz_test.f1
	/src/github.com/pinzolo/errz/errz_test.go:19
github.com/pinzolo/errz_test.f
	/src/github.com/pinzolo/errz/errz_test.go:12
wrap1
github.com/pinzolo/errz_test.TestFormat.func2.1
	/src/github.com/pinzolo/errz/errz_test.go:96
github.com/pinzolo/errz_test.f3
	/src/github.com/pinzolo/errz/errz_test.go:31
github.com/pinzolo/errz_test.f2
	/src/github.com/pinzolo/errz/errz_test.go:26
github.com/pinzolo/errz_test.f1
	/src/github.com/pinzolo/errz/errz_test.go:19
github.com/pinzolo/errz_test.f
	/src/github.com/pinzolo/errz/errz_test.go:12
wrap2
github.com/pinzolo/errz_test.TestFormat.func2
	/src/github.com/pinzolo/errz/errz_test.go:95
github.com/pinzolo/errz_test.f2
	/src/github.com/pinzolo/errz/errz_test.go:26
github.com/pinzolo/errz_test.f1
	/src/github.com/pinzolo/errz/errz_test.go:19
github.com/pinzolo/errz_test.f
	/src/github.com/pinzolo/errz/errz_test.go:12
`
	if buf.String() != want {
		t.Error("invalid format result")
		t.Errorf("got: %s", buf.String())
		t.Errorf("want: %s", want)
	}
}

func newStack(cs []*errz.Caller, more bool) *errz.StackTrace {
	return &errz.StackTrace{
		More:    more,
		Callers: cs,
	}
}

func callers1() []*errz.Caller {
	return []*errz.Caller{
		newCaller("github.com/pinzolo/errz_test", "TestFormat.func2.1.1", "/src/github.com/pinzolo/errz/errz_test.go", 97),
		newCaller("github.com/pinzolo/errz_test", "f4", "/src/github.com/pinzolo/errz/errz_test.go", 38),
		newCaller("github.com/pinzolo/errz_test", "f3", "/src/github.com/pinzolo/errz/errz_test.go", 33),
		newCaller("github.com/pinzolo/errz_test", "f2", "/src/github.com/pinzolo/errz/errz_test.go", 24),
		newCaller("github.com/pinzolo/errz_test", "f1", "/src/github.com/pinzolo/errz/errz_test.go", 19),
		newCaller("github.com/pinzolo/errz_test", "f", "/src/github.com/pinzolo/errz/errz_test.go", 12),
	}
}

func callers2() []*errz.Caller {
	return []*errz.Caller{
		newCaller("github.com/pinzolo/errz_test", "TestFormat.func2.1", "/src/github.com/pinzolo/errz/errz_test.go", 96),
		newCaller("github.com/pinzolo/errz_test", "f3", "/src/github.com/pinzolo/errz/errz_test.go", 31),
		newCaller("github.com/pinzolo/errz_test", "f2", "/src/github.com/pinzolo/errz/errz_test.go", 26),
		newCaller("github.com/pinzolo/errz_test", "f1", "/src/github.com/pinzolo/errz/errz_test.go", 19),
		newCaller("github.com/pinzolo/errz_test", "f", "/src/github.com/pinzolo/errz/errz_test.go", 12),
	}
}

func callers3() []*errz.Caller {
	return []*errz.Caller{
		newCaller("github.com/pinzolo/errz_test", "TestFormat.func2", "/src/github.com/pinzolo/errz/errz_test.go", 95),
		newCaller("github.com/pinzolo/errz_test", "f2", "/src/github.com/pinzolo/errz/errz_test.go", 26),
		newCaller("github.com/pinzolo/errz_test", "f1", "/src/github.com/pinzolo/errz/errz_test.go", 19),
		newCaller("github.com/pinzolo/errz_test", "f", "/src/github.com/pinzolo/errz/errz_test.go", 12),
	}
}

func newCaller(pkg, fn, file string, line int) *errz.Caller {
	return &errz.Caller{
		Package: pkg,
		Func:    fn,
		File:    file,
		Line:    line,
	}
}
