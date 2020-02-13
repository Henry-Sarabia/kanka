package kanka

import (
	"fmt"
	"net/http"
)

// serverError represents an error originating from another server.
type serverError struct {
	code      int
	status    string
	temporary bool
}

// Error returns the status message of an error.
func (e *serverError) Error() string {
	return fmt.Sprintf("server responded with status '%s'", e.status)
}

// Temporary returns true if the error is temporary.
func (e *serverError) Temporary() bool {
	return e.temporary
}

// isSuccess returns true if the provided status code is of the 200 type.
func isSuccess(code int) bool {
	if code >= 200 && code < 300 {
		return true
	}

	return false
}

// isTemporary returns true if the provided status code represents a temporary
// error according to Kanka.
// For more information, visit: https://kanka.io/en-US/docs/1.0/setup#endpoints
func isTemporary(code int) bool {
	switch code {
	case http.StatusMisdirectedRequest:
		return true
	case http.StatusTooManyRequests:
		return true
	default:
		return false
	}
}
