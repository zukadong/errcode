package errcode

type (
	Error interface {
		Error() string
		WithCode(int) Error
		Code() int
		WithMessage(string) Error
		Message() string
		WithArgs(...any) Error
		Args() []any
	}

	defaultError struct {
		err  error
		C    int    `json:"code"`
		M    string `json:"message"`
		args []any
	}
)

func Wrap(err error) Error {
	return &defaultError{
		err: err,
	}
}

func (e *defaultError) Error() string {
	return e.err.Error()
}

func (e *defaultError) WithCode(code int) Error {
	e.C = code
	return e
}

func (e *defaultError) Code() int {
	return e.C
}

func (e *defaultError) WithMessage(message string) Error {
	e.M = message
	return e
}

func (e *defaultError) Message() string {
	return e.M
}

func (e *defaultError) WithArgs(args ...any) Error {
	e.args = append(e.args, args...)
	return e
}

func (e *defaultError) Args() []any {
	return e.args
}
