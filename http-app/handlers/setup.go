package handlers

import (
	"net/http"
)

func SetupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/healthcheck", healthCheckHandler)
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/", indexHandler)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

}
