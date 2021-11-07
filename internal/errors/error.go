package errors

import "github.com/go-ole/go-ole"

// ComError is implemented by both Error and ole.OleError.
type ComError interface {
	error
	Code() uintptr
	SubError() error
}

type Error struct {
	code        uintptr
	description string
	err         error
}

func (e *Error) Code() uintptr {
	return e.code
}

func (e *Error) Error() string {
	return e.description
}

func (e *Error) SubError() error {
	return e.err
}

func NotImplemented(err error) *Error {
	return &Error{
		code: ole.E_NOTIMPL,
		err:  err,

		// TODO: Consider localizing the description for Windows like ole.OleError.
		description: "Not implemented",
	}
}
