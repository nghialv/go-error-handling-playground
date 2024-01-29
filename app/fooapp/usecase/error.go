package usecase

import "github.com/nghialv/go-error-handling-playground/pkg/cerror"

var (
	ErrNotFound = cerror.New("not found")
)
