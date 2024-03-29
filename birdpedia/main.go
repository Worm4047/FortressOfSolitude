package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Declare a new router
	r := mux.NewRouter()

	r.HandleFunc("/hello", handler).Methods("GET")

	http.ListenAndServe(":8080", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
