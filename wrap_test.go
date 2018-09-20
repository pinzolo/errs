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
		t.Errorf("should have stack trace")
	}
}

func TestWrapWithNil(t *testing.T) {
	if errz.Wrap(nil, "wrap") != nil {
		t.Errorf("Wrap should return nil when given nil")
	}
}

func TestWrapf(t *testing.T) {
	orig := errors.New("original")
	err := errz.Wrapf(orig, "wrap %d %s %q", 1, "2", "3")

	if err.Error() != "wrap 1 2 \"3\": original" {
		t.Errorf("invalid error message: want %q, got %q", "wrap: original", err.Error())
	}

	if errz.Trace(err) == nil {
		t.Errorf("should have stack trace")
	}
}

func TestWrapfWithNil(t *testing.T) {
	if errz.Wrapf(nil, "wrap %d %s %q", 1, "2", "3") != nil {
		t.Errorf("Wrapf should return nil when given nil")
	}
}

func TestCause(t *testing.T) {
	orig := errors.New("original")
	data := []struct {
		err  error
		want error
		memo string
	}{
		{nil, nil, "nil"},
		{orig, orig, "raw error"},
		{errz.Wrap(orig, "wrap"), orig, "wrap"},
		{errz.Wrap(errz.Wrap(errz.Wrap(orig, "wrap1"), "wrap2"), "wrap3"), orig, "nested wrap"},
	}

	for _, d := range data {
		t.Run(d.memo, func(t *testing.T) {
			err := errz.Cause(d.err)
			if err != d.want {
				t.Errorf("cause error: want %v, got %v", d.want, err)
			}
		})
	}
}

func TestIsWrapped(t *testing.T) {
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
			if errz.IsWrapped(d.err) != d.want {
				if d.want {
					t.Error("IsWrapped should return true if error is wrapped")
				} else {
					t.Error("IsWrapped should return false if error is not wrapped")
				}
			}
		})
	}
}
