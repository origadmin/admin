//go:build !embed_ui

package web

import (
	"errors"
	"net/http"
)

// GetHandler is a stub implementation for when the UI is not embedded.
// It returns an error, indicating that the embedded UI is not available.
// This version is compiled by default, unless the 'embed_ui' build tag is provided.
func GetHandler() (http.Handler, error) {
	return nil, errors.New("web UI is not embedded in this build")
}
