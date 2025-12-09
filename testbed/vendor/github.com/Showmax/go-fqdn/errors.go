package fqdn

import "fmt"

// Error for cases when os.Hostname() fails.
var ErrHostnameFailed = errHostnameFailed{}

// Error for cases when we could not found fqdn for whatever reason.
var ErrFqdnNotFound = errFqdnNotFound{}

type errHostnameFailed struct {
	cause error
}

func (e errHostnameFailed) Error() string {
	return fmt.Sprintf("could not get hostname: %v", e.cause)
}

func (e errHostnameFailed) Unwrap() error {
	return e.cause
}

func (e errHostnameFailed) Is(target error) bool {
	switch target.(type) {
	case errHostnameFailed:
		return true
	default:
		return false
	}
}

type errFqdnNotFound struct {
	cause error
}

func (e errFqdnNotFound) Error() string {
	return fmt.Sprintf("fqdn hostname not found: %v", e.cause)
}

func (e errFqdnNotFound) Unwrap() error {
	return e.cause
}

func (e errFqdnNotFound) Is(target error) bool {
	switch target.(type) {
	case errFqdnNotFound:
		return true
	default:
		return false
	}
}
