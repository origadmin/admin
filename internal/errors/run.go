package errors

import (
	"fmt"
)

type runError struct {
	name    string
	message string
	err     error
}

func (r runError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", r.name, r.message, r.err)
}

func RunError(name string, message string, err error) error {
	return &runError{
		name:    name,
		message: message,
		err:     err,
	}
}
