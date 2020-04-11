package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	http.HandleFunc("/sha", shaHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func shaHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "hello GET"}`))
	case "POST":
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "hello POST"}`))
	}
}
