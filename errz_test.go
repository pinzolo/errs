package errz_test

import (
	"testing"

	"github.com/pinzolo/errz"
)

func f(i int, raise func() error) error {
	return f1(i, raise)
}

func f1(i int, raise func() error) error {
	if i == 1 {
		return raise()
	}
	return f2(i, raise)
}

func f2(i int, raise func() error) error {
	if i == 2 {
		return raise()
	}
	return f3(i, raise)
}

func f3(i int, raise func() error) error {
	if i == 3 {
		return raise()
	}
	return f4(i, raise)
}

func f4(i int, raise func() error) error {
	if i == 4 {
		return raise()
	}
	return f5(i, raise)
}

func f5(i int, raise func() error) error {
	if i == 5 {
		return raise()
	}
	return nil
}

func TestNew(t *testing.T) {
	err := errz.New("new error")

	if err.Error() != "new error" {
		t.Errorf("invalid error message: want %q, got %q", "new error", err.Error())
	}

	if errz.Trace(err) == nil {
		t.Errorf("should has stack trace")
	}
}

func TestErrorf(t *testing.T) {
	err := errz.Errorf("error %s", "test")

	if err.Error() != "error test" {
		t.Errorf("invalid error message: want %q, got %q", "error test", err.Error())
	}

	if errz.Trace(err) == nil {
		t.Errorf("should has stack trace")
	}
}
