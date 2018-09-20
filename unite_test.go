package errz_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/pinzolo/errz"
)

func TestUnite(t *testing.T) {
	err1 := errors.New("error1")
	err2 := errors.New("error2")
	err3 := errors.New("error3")
	data := []struct {
		err1 error
		err2 error
		msg  string
		memo string
	}{
		{nil, nil, "", "both nil"},
		{err1, nil, "error1", "nil err2"},
		{nil, err2, "error2", "nil err1"},
		{err1, err2, "contains 2 errors...\n1) error1\n2) error2\n", "unite"},
		{errz.Unite(err1, err2), err3, "contains 3 errors...\n1) error1\n2) error2\n3) error3\n", "nested"},
	}
	for _, d := range data {
		t.Run(d.memo, func(t *testing.T) {
			err := errz.Unite(d.err1, d.err2)
			if d.err1 == nil && d.err2 == nil {
				if err != nil {
					t.Fatal("should return nil when both errors are nil")
				}
				return
			}
			if err == nil {
				t.Fatal("Unite should not return nil when either error is not nil")
			}
			if err.Error() != d.msg {
				t.Errorf("Unite invalid error message: want %s, got %s", d.msg, err.Error())
			}
		})
	}
}

func TestErrs(t *testing.T) {
	err1 := errors.New("error1")
	err2 := errors.New("error2")
	err3 := errors.New("error3")
	data := []struct {
		err1 error
		err2 error
		errs []error
		memo string
	}{
		{nil, nil, nil, "both nil"},
		{err1, nil, []error{err1}, "nil err2"},
		{nil, err2, []error{err2}, "nil err1"},
		{err1, err2, []error{err1, err2}, "unite"},
		{errz.Unite(err1, err2), err3, []error{err1, err2, err3}, "nested"},
	}
	for _, d := range data {
		t.Run(d.memo, func(t *testing.T) {
			err := errz.Unite(d.err1, d.err2)
			if d.err1 == nil && d.err2 == nil {
				if errz.Errs(err) != nil {
					t.Fatal("Errs should return nil when both errors are nil")
				}
				return
			}
			if err == nil {
				t.Fatal("Errs should not return nil when either error is not nil")
			}
			if !reflect.DeepEqual(errz.Errs(err), d.errs) {
				t.Errorf("invalid error message: want %+v, got %+vs", d.errs, errz.Errs(err))
			}
		})
	}
}
