package cerror

import (
	"errors"
	"fmt"
)

type Error struct {
	error
	stack       *stack
	reasonCodes []ReasonCode
}

func (e *Error) StackTrace() string {
	if e.stack == nil {
		return ""
	}

	return fmt.Sprintf("%+v", e.stack)
}

func (e *Error) Unwrap() error {
	return e.error
}

func New(msg string, opts ...Option) error {
	err := &Error{
		error: fmt.Errorf(msg),
		stack: callers(),
	}

	for _, opt := range opts {
		opt(err)
	}

	return err
}

func Wrap(err error, prefix string, opts ...Option) error {
	werr := &Error{
		error: fmt.Errorf("%s: %w", prefix, err),
	}

	if cerr, ok := As(err); ok {
		werr.stack = cerr.stack
	} else {
		werr.stack = callers()
	}

	for _, opt := range opts {
		opt(werr)
	}

	return werr
}

func As(err error) (*Error, bool) {
	cerr := &Error{}
	if ok := errors.As(err, &cerr); !ok {
		return nil, false
	}

	return cerr, true
}
