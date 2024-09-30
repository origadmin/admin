package errors

import (
	"fmt"
)

type configErr struct {
	name    string
	message string
	error
}

func (c configErr) Error() string {
	return fmt.Sprintf("file %s: %s,error:%v", c.name, c.message, c.error)
}

func LoadError(err error, name string, message string) error {
	return configErr{name: name, message: message, error: err}
}
