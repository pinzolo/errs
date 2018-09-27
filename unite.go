package errz

import (
	"bytes"
	"fmt"
)

type withErrors struct {
	errs []error
}

// nolint: gas
func (we *withErrors) Error() string {
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "contains %d errors...\n", len(we.errs))
	for i, err := range we.errs {
		fmt.Fprintf(buf, "%d) %s\n", i+1, err.Error())
	}
	return buf.String()
}

func (we *withErrors) Errs() []error {
	return we.errs
}

// Unite 2 errors to 1 error that having internal error list.
// If both error are nil, Unite returns nil.
// If either error is nil, Unite returns another error.
func Unite(err1, err2 error) error {
	if err1 == nil && err2 == nil {
		return nil
	}
	if err1 == nil {
		return err2
	}
	if err2 == nil {
		return err1
	}

	errs1 := Errs(err1)
	errs2 := Errs(err2)
	errs := make([]error, len(errs1)+len(errs2))
	for i, e := range errs1 {
		errs[i] = e
	}
	for i, e := range errs2 {
		errs[len(errs1)+i] = e
	}
	return &withErrors{errs}
}

// Errs returns error list in given error.
// If given error is nil, Errs returns nil.
// If given error does not have error list, Errs returns error list that contains given error itself only.
func Errs(err error) []error {
	if err == nil {
		return nil
	}

	type errser interface {
		Errs() []error
	}
	if errs, ok := err.(errser); ok {
		return errs.Errs()
	}
	return []error{err}
}
