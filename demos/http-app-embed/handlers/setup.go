package handlers

import (
	"embed"
	"net/http"
)

//go:embed static
var staticFS embed.FS

// staticFS
// - css/styles.css
// - js/index.js

func SetupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/healthcheck", healthCheckHandler)
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/", indexHandler)

	staticFileServer := http.FileServer(http.FS(staticFS))
	mux.Handle("/static/", staticFileServer)
}
