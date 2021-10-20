package main

import (
	"github.com/amitsaha/project1/milestone2-code/webapp/handlers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handlers.SetupHandlers(mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
