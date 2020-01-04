package kanka

import (
	"fmt"
	"net/http"
)

// ServerError represents a
type ServerError struct {
	code      int
	status    string
	temporary bool
}

func (e *ServerError) Error() string {
	return fmt.Sprintf("server responded with status '%s'", e.status)
}

// Temporary returns true if the error is temporary.
func (e *ServerError) Temporary() bool {
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
func isTemporary(code int) bool {
	switch code {
	case http.StatusMisdirectedRequest:
		return true
	default:
		return false
	}
}
