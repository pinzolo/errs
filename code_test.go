package errz_test

import (
	"errors"
	"testing"

	"github.com/pinzolo/errz"
)

func TestWithCode(t *testing.T) {
	data := []struct {
		err  error
		code string
		msg  string
		note string
	}{
		{errors.New("error1"), "E01", "[E01] error1", "valid"},
		{nil, "E02", "", "nil error"},
		{errors.New("error3"), "", "error3", "empty code"},
	}

	for _, d := range data {
		t.Run(d.note, func(t *testing.T) {
			err := errz.WithCode(d.err, d.code)
			if d.err == nil {
				if err != nil {
					t.Errorf("WithCode should return nil when given error is nil")
				}
				if errz.Code(err) != "" {
					t.Errorf("Code should return empty when given error is nil")
				}
			} else {
				if err == nil {
					t.Errorf("WithCode should not return nil when given error is not nil")
				}
				if err.Error() != d.msg {
					t.Errorf("error message: want %s, got %s", d.msg, err.Error())
				}
				if cd := errz.Code(err); cd != d.code {
					t.Errorf("code extract failure: want %q, got %q", d.code, cd)
				}
			}
		})
	}
}
