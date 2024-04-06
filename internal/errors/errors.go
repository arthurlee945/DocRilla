package errors

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

const (
	ErrUnknown        = Error("err_unknwon: unknown error occured")
	ErrInvalidRequest = Error("err_invalid_request: invalid request received")
	ErrValidation     = Error("err_validation: failed validation")
	ErrNotFound       = Error("err_not_found: not found")
)

const ErrSeperator = " -- "

/*
built-in errors pkg error interface - for reference
type error interface {
    Error() string
}
*/

type Error string

func (e Error) Error() string {
	return string(e)
}

func (e Error) Is(target error) bool {
	return e.Error() == target.Error() || strings.HasPrefix(target.Error(), e.Error()+ErrSeperator)
}

func (e Error) As(target interface{}) bool {
	if v := reflect.ValueOf(target).Elem(); v.Type().Name() == "Error" && v.CanSet() {
		v.SetString(e.Error())
		return true
	}
	return false
}

func (e Error) Wrap(err error) wrappedError {
	return wrappedError{msg: e.Error(), cause: err}
}

type wrappedError struct {
	msg   string
	cause error
}

func (w wrappedError) Error() string {
	if w.cause != nil {
		return fmt.Sprintf("%s%s%v", w.msg, ErrSeperator, w.cause)
	}
	return w.msg
}

func (w wrappedError) Is(target error) bool {
	return Error(w.msg).Is(target)
}

func (w wrappedError) As(target interface{}) bool {
	return Error(w.msg).As(target)
}

func (w wrappedError) Unwrap() error {
	return w.cause
}

// utility syntactic sugar for built-in errors pkg
func New(msg string) error {
	return errors.New(msg)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, target)
}
