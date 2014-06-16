package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func logRequest(
	handler func(http.ResponseWriter, *http.Request),
) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s\n", r.Method, r.URL.String())
		handler(w, r)
	}
}

func MakeRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", logRequest(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world!")
	}))

	return r
}
