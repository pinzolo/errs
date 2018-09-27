package errz

import "fmt"

// Const is able to declare error as constant.
// e.g.
// const EOF = errz.Const("EOF")
type Const string

// Error is error message of Const.
func (c Const) Error() string {
	return string(c)
}

// Errorf builds error message with given args by using itself as format.
func (c Const) Errorf(a ...interface{}) string {
	return fmt.Sprintf(string(c), a...)
}

// WithTrace returns new error with stack trace based Const.
func (c Const) WithTrace() error {
	return &wrapper{
		err: c,
		pcs: pcs(defaultSkip),
	}
}
