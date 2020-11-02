package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// topLevel is a handler for displaying the welcome screen.
func topLevel(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "Welcome!\n\nAppend'/x' to the URL (where x is a number 1-4), to enable concurrent threads/Goroutines.")
}

// getThreads is a handler for reciving user input from the URL.
func getThreads(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	// vars is a call to the mux router to recive the varibles from the http request.
	vars := mux.Vars(r)

	// threads is a string from the list of varibles provided by mux.Vars.
	threads := vars["Id=threads"]

	// intThreads is the varible "threads" converted fromstring to int.
	intThreads, err := strconv.Atoi(threads)
	if err != nil {
		fmt.Fprintln(w, "Invalid input ("+threads+").")
		return
	}
	if intThreads == 0 {
		fmt.Fprintln(w, "Threads cannot be 0")
		return
	}
	if intThreads > len(urls) {
		fmt.Fprintln(w, "Threads ("+threads+") exceedes number of URLS ("+fmt.Sprintf("%d", len(urls))+").\n")
		return
	}

	fmt.Fprintln(w, "Threads: "+threads+".\n")

	titles, status := getTitle(intThreads)

	fmt.Fprintln(w, fmt.Sprintf("%d", status[0])+" titles were found:\n")
	for i := 0; i < len(titles); i++ {
		fmt.Fprintln(w, titles[i])
	}
	fmt.Fprintln(w, "\nThe number of successful calls were: "+fmt.Sprintf("%d", status[0])+".")
	fmt.Fprintln(w, "The number of failed calls were: "+fmt.Sprintf("%d", status[1])+".")
	return
}
