package errz

type withCode struct {
	error
	code string
}

func (wc *withCode) Error() string {
	return "[" + wc.code + "] " + wc.error.Error()
}

func (wc *withCode) Code() string {
	return wc.code
}

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
	return &withCode{
		error: err,
		code:  code,
	}
}

// Code returns error code that given error annotated with.
// If given error is nil, Code returns empty.
// If given error is raw error, Code returns empty,
func Code(err error) string {
	if err == nil {
		return ""
	}

	type coder interface {
		Code() string
	}

	if c, ok := err.(coder); ok {
		return c.Code()
	}
	return ""
}
