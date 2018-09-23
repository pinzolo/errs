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
	if b, ok := err.(*box); ok {
		b.reset(defaultSkip)
		b.code = code
		return b
	}

	return &box{
		code:  code,
		msg:   err.Error(),
		cause: newCause(err, defaultSkip),
	}
}

// Code returns error code that given error annotated with.
// If given error is nil, Code returns empty.
// If given error is raw error, Code returns empty,
func Code(err error) string {
	if err == nil {
		return ""
	}

	if b, ok := err.(*box); ok {
		return b.Code()
	}
	return ""
}
