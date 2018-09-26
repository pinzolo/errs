package errz

// NewWrapper creates custom errz.wrapper instance for testing.
func NewWrapper(err error, code string, msg string, st *StackTrace) error {
	return &wrapper{
		err:  err,
		code: code,
		msg:  msg,
		st:   st,
	}
}
