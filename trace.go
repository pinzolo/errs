package errz

// Trace returns stack trace.
// If err is nil, returns nil.
func Trace(err error) *StackTrace {
	if err == nil {
		return nil
	}

	type tracer interface {
		Trace() *StackTrace
	}

	if wt, ok := err.(tracer); ok {
		return wt.Trace()
	}
	return nil
}
