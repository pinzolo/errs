package errz

import "fmt"

// Wrap returns new error annotated given message.
// Returned error has stack trace and original error.
// If cause error is nil, Wrap returns nil.
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	msg := newMsg(err, message)

	if b, ok := err.(*box); ok {
		b.reset(defaultSkip)
		b.msg = msg
		return b
	}
	return &box{
		trace: newTrace(defaultSkip),
		cause: newCause(err, defaultSkip),
		msg:   msg,
	}
}

// Wrapf returns new error annotated message that is build given format and args.
// Returned error has stack trace and original error.
// If cause error is nil, Wrap returns nil.
func Wrapf(err error, format string, a ...interface{}) error {
	if err == nil {
		return nil
	}
	msg := newMsg(err, fmt.Sprintf(format, a...))

	if b, ok := err.(*box); ok {
		b.reset(defaultSkip)
		b.msg = msg
		return b
	}
	return &box{
		trace: newTrace(defaultSkip),
		cause: newCause(err, defaultSkip),
		msg:   msg,
	}
}

func newMsg(err error, message string) string {
	msg := err.Error()
	if message != "" {
		msg = message + ": " + msg
	}
	return msg
}

// Cause returns original error that is wrapped by errz.Wrap or errz.Wrapf recursively.
// If given error is not wrapped, Cause returns itself.
// If given error is nil, Cause returns nil.
func Cause(err error) error {
	if err == nil {
		return nil
	}

	if b, ok := err.(*box); ok {
		return b.Cause()
	}
	return err
}
