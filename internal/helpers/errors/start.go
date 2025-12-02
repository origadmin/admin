/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package errors

import (
	"fmt"
)

type startError struct {
	name    string
	message string
	err     error
}

func (r startError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", r.name, r.message, r.err)
}

func StartError(name string, message string, err error) error {
	return &startError{
		name:    name,
		message: message,
		err:     err,
	}
}
