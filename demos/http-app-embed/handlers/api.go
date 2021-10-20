package handlers

import "net/http"
import "fmt"

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have reached Echorand Corp's Service API")
}
