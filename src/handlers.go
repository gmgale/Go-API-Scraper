package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	"github.com/gorilla/mux"
)

// topLevel is a handler for displaying the welcome screen.
func topLevel(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintln(w, "Welcome!\n\nAppend'/x' to the URL (where x is a number 1-4), to enable concurrent threads/Goroutines.")
}

// shutdownHandler is a handler for starting API shutdown request
func (s *myServer) shutdownHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Shutdown server"))

	// Do nothing if shutdown request already issued
	// if s.reqCount == 0 then set to 1, return true otherwise false
	if !atomic.CompareAndSwapUint32(&s.reqCount, 0, 1) {
		log.Printf("Shutdown through API call in progress...")
		return
	}

	go func() {
		s.shutdownReq <- true
	}()
}

// getThreads is a handler for reciving user input from the URL.
func getThreads(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	// clockstart is used later for mesuring the duration of the process.
	clockStart := time.Now()

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

	newData := getTitle(intThreads)

	fmt.Fprintln(w, fmt.Sprintf("%d", newData.status.succeeded)+" titles were found:\n")
	for i := 0; i < len(newData.results); i++ {
		fmt.Fprintln(w, newData.results[i].url)
		fmt.Fprintln(w, newData.results[i].title)
	}
	fmt.Fprintln(w, "\nThe number of successful calls were: "+fmt.Sprintf("%d", newData.status.succeeded)+".")
	fmt.Fprintln(w, "The number of failed calls were: "+fmt.Sprintf("%d", newData.status.failed)+".")

	clockStop := time.Now()

	// finish populating newData to be sent via JSON to database
	newData.time = clockStart.String()
	newData.duration = clockStop.Sub(clockStart).String()
	newData.threads = intThreads
	newData.id = globalCallCounter

	globalCallCounter++

	dbSend(newData)
	return
}
