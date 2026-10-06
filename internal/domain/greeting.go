package domain

import "errors"

var ErrEmptyName = errors.New("name must not be empty")

// Greeting is the core concept returned by the hello use case.
type Greeting struct {
	Message string
}
