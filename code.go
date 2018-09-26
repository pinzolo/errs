package errz

// WithCode annotates given erro with given code.
// If given error is nil, WithCode returns nil.
// If given code is empty, WithCode returns given error itself.
func WithCode(err error, code string) error {
	if err == nil {
		return nil
	}
	if code == "" {
		return err
	}

	return &wrapper{
		err:  err,
		code: code,
		pcs:  pcs(defaultSkip),
	}
}

// Code returns error code that given error annotated with.
// If given error is nil, Code returns empty.
// If given error is raw error, Code returns empty,
func Code(err error) string {
	if err == nil {
		return ""
	}

	if w, ok := err.(*wrapper); ok {
		return w.code
	}
	return ""
}
