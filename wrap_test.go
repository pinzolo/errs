package errz_test

import (
	"errors"
	"testing"

	"github.com/pinzolo/errz"
)

func TestWrap(t *testing.T) {
	orig := errors.New("original")
	err := errz.Wrap(orig, "wrap")

	if err.Error() != "wrap: original" {
		t.Errorf("invalid error message: want %q, got %q", "wrap: original", err.Error())
	}

	if errz.Trace(err) == nil {
		t.Errorf("should has stack trace")
	}
}

func TestWrapf(t *testing.T) {
	orig := errors.New("original")
	err := errz.Wrapf(orig, "wrap %d %s %q", 1, "2", "3")

	if err.Error() != "wrap 1 2 \"3\": original" {
		t.Errorf("invalid error message: want %q, got %q", "wrap: original", err.Error())
	}

	if errz.Trace(err) == nil {
		t.Errorf("should has stack trace")
	}
}
