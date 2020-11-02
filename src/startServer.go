package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// startServer sets up a localhost server using the gorilla/mux package
// and calls handlers for endpoints.
func startServer() {
	log.Println("startServer funtion called.")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api", topLevel)
	router.HandleFunc("/api/{Id=threads}", getThreads)
	log.Fatal(http.ListenAndServe(":8080", router))
}
