package errz

type trace struct {
	pcs []uintptr
	st  *StackTrace
}

func (t *trace) Trace() *StackTrace {
	if t.st != nil {
		return t.st
	}

	t.st = &StackTrace{}
	cs := callers(t.pcs)
	if len(cs) > MaxDepth {
		t.st.More = true
		cs = cs[0:MaxDepth]
	}
	t.st.Callers = cs
	return t.st
}

func newTrace(skip int) *trace {
	return &trace{
		pcs: pcs(skip + 1),
	}
}

type cause struct {
	*trace
	err error
}

func (c *cause) Error() string {
	return c.err.Error()
}

func newCause(err error, skip int) *cause {
	return &cause{
		trace: newTrace(skip + 1),
		err:   err,
	}
}

type box struct {
	*trace
	*cause
	msg  string
	code string
}

func (b *box) Error() string {
	if b.code == "" {
		return b.msg
	}
	return "[" + b.code + "] " + b.msg
}

func (b *box) Code() string {
	return b.code
}

func (b *box) Cause() error {
	return b.cause.err
}

func (b *box) CauseTrace() *StackTrace {
	return b.cause.Trace()
}

func (b *box) reset(skipTrace int) {
	b.st = nil
	b.pcs = pcs(skipTrace + 1)
}
