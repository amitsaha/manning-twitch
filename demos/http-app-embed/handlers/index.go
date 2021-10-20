package handlers

import (
	_ "embed"
	"fmt"
	"net/http"
)

//go:embed index.html
var indexHtml string

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, indexHtml)
}
