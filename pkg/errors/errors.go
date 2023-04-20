package errors

import (
	"errors"
	"fmt"
)

var (
	ErrFailedLoad    = errors.New("unable to load model")
	ErrFailedExec    = errors.New("unable to execute model")
	ErrInvalidResult = errors.New("invalid result")
)

func NewFailedExec(msg any) error {
	return fmt.Errorf("%w: %v", ErrFailedExec, msg)
}

func NewFailedLoad(msg any) error {
	return fmt.Errorf("%w: %v", ErrFailedLoad, msg)
}

func NewInvalidResult(msg any) error {
	return fmt.Errorf("%w: %v", ErrInvalidResult, msg)
}
