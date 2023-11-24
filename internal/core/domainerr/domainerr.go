package domainerr

import (
	"errors"
	"fmt"
)

var (
	errRoot = errors.New("domain defined error")
)

func New(msg string) error {
	return fmt.Errorf("%w: %v", errRoot, msg)
}

func Is(err error) bool {
	return errors.Is(err, errRoot)
}

func Join(errs ...error) error {
	return errors.Join(errs...)
}
