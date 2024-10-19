package app_error

import "fmt"

// HttpError is an error that can be returned to the client.
// It contains an error message, an http status code, and the root cause error.
type HttpError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e HttpError) Error() string {
	return fmt.Sprintf("%d: %s", e.StatusCode, e.Message)
}
