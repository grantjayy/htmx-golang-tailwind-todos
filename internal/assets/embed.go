package assets

import (
	"embed"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// The "//go:embed all:dist" directive tells the Go compiler to embed the contents of the dist directory
// into the Assets variable. The Assets variable is of type embed.FS, which is a file system interface
// that can be used like a regular file system.

//go:embed all:static
var Assets embed.FS

// ServeAssets mount the embedded assets to an HTTP server
func Mount(r chi.Router) {
	r.Route("/static", func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				w.Header().Add("Link", fmt.Sprintf("<%s>; rel=preload;", r.URL.Path))

				if strings.HasSuffix(r.URL.Path, ".js") {
					w.Header().Add("Content-Type", "application/javascript")
					w.Header().Add("Content-Encoding", "br")
				} else if strings.HasSuffix(r.URL.Path, ".css") {
					w.Header().Add("Content-Type", "text/css")
					w.Header().Add("Content-Encoding", "br")
				}
				next.ServeHTTP(w, r)
			})
		})
		r.Handle("/*", http.FileServer(http.FS(Assets)))
	})
}
