package error

import "errors"

type Code int

const (
	InvalidArgument Code = iota
	NotFound
	Unauthorized
)

var (
	ErrNotFound = NewError(NotFound, "対象のデータが見つかりませんでした")
)

type Error struct {
	code    Code
	message string
}

func (e *Error) Error() string {
	return e.message
}

func (e *Error) Is(target error) bool {
	var t *Error
	if ok := errors.As(target, &t); !ok {
		return false
	}
	return e.code == t.code
}

func NewError(c Code, m string) *Error {
	return &Error{
		code:    c,
		message: m,
	}
}
