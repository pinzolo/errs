package errz

import "fmt"

// Wrap returns new error annotated given message.
// Returned error has stack trace and original error.
// If cause error is nil, Wrap returns nil.
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}

	return &wrapper{
		err: err,
		msg: message,
		pcs: pcs(defaultSkip),
	}
}

// Wrapf returns new error annotated message that is build given format and args.
// Returned error has stack trace and original error.
// If cause error is nil, Wrap returns nil.
func Wrapf(err error, format string, a ...interface{}) error {
	if err == nil {
		return nil
	}

	return &wrapper{
		err: err,
		msg: fmt.Sprintf(format, a...),
		pcs: pcs(defaultSkip),
	}
}

// Cause returns original error that is wrapped by errz.Wrap or errz.Wrapf recursively.
// If given error is not wrapped, Cause returns itself.
// If given error is nil, Cause returns nil.
func Cause(err error) error {
	if err == nil {
		return nil
	}

	if w, ok := err.(*wrapper); ok {
		if w.err == nil {
			return w
		}
		return Cause(w.err)
	}
	return err
}
