package handlers

import (
	"io"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "ok")
}
