package errz_test

import (
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
