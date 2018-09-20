package errz

import "fmt"

type withCause struct {
	*base
	cause error
}

func (wc *withCause) Cause() error {
	return wc.cause
}

func (wc *withCause) Error() string {
	return wc.msg + ": " + wc.cause.Error()
}

// Wrap returns new error annotated given message.
// Returned error has stack trace and original error.
// If cause error is nil, Wrap returns nil.
func Wrap(cause error, msg string) error {
	if cause == nil {
		return nil
	}
	return &withCause{
		cause: cause,
		base: &base{
			msg: msg,
			st:  callers(),
		},
	}
}

// Wrapf returns new error annotated message that is build given format and args.
// Returned error has stack trace and original error.
// If cause error is nil, Wrap returns nil.
func Wrapf(cause error, format string, a ...interface{}) error {
	if cause == nil {
		return nil
	}
	return &withCause{
		cause: cause,
		base: &base{
			msg: fmt.Sprintf(format, a...),
			st:  callers(),
		},
	}
}

type causer interface {
	Cause() error
}

// Cause returns original error that is wrapped by errz.Wrap or errz.Wrapf recursively.
// If given error is not wrapped, Cause returns itself.
// If given error is nil, Cause returns nil.
func Cause(err error) error {
	if err == nil {
		return nil
	}

	if c, ok := err.(causer); ok {
		return Cause(c.Cause())
	}
	return err
}

// IsWrapped returns true if given error is already wrapped.
// If given error is nil, IsWrapped returns false,
func IsWrapped(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(causer)
	return ok
}
