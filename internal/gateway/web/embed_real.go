//go:build embed_ui

package web

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed all:../../../../resources/web
var WebUI embed.FS

// GetHandler returns an http.Handler that serves the embedded Web UI.
// This version is compiled only when the 'embed_ui' build tag is provided.
func GetHandler() (http.Handler, error) {
	// The `WebUI` embed.FS now contains the `resources/web` directory structure.
	// We need to create a sub-filesystem that starts from that directory.
	distFS, err := fs.Sub(WebUI, "resources/web")
	if err != nil {
		return nil, err
	}

	// Create a file server for the sub-filesystem.
	fileServer := http.FileServer(http.FS(distFS))

	// Create a handler that serves static files and falls back to index.html for SPAs.
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Try to serve the static file first.
		_, err := distFS.Open(r.URL.Path[1:])
		if err == nil {
			fileServer.ServeHTTP(w, r)
			return
		}

		// If the file is not found, serve index.html.
		r.URL.Path = "/"
		fileServer.ServeHTTP(w, r)
	})

	return handler, nil
}
