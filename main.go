package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

func main() {
	startServer()
}

// Set up a localhost server
func startServer() {
	log.Println("startServer funtion called.")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api", topLevel)
	router.HandleFunc("/api/{Id=threads}", getThreads)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Welcome displays splash screen at .../api
func topLevel(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	
	fmt.Fprintln(w, "Welcome!\n\nAppend'/x' to the URL (where x is a number 1-4), to enable concurrent threads/goroutines.")
}

// API Threads Endpoint ---------------------------------------------------------
func getThreads(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	urls := []string{
		"https://www.result.si/projekti/",
		"https://www.result.si/o-nas/",
		"https://www.result.si/kariera/",
		"https://www.result.si/blog/"}

	// Make channels for Go Routines
	// urlChannel is for title data
	// statChannel is for GET url succ/fail count
	urlCh := make(chan string, 4)
	statCh := make(chan string, 4)
	defer close(urlCh)
	defer close(statCh)

	// Get threads from URL
	vars := mux.Vars(r)
	threads := vars["Id=threads"]

	// Convert threads string to int
	intThreads, err := strconv.Atoi(threads)
	if err != nil {
		// Handle error
		fmt.Fprintln(w, "Invalid input ("+threads+").")
		return
	}

	if intThreads > len(urls) {
		fmt.Fprintln(w, "Threads ("+threads+") exceedes number of URLS ("+fmt.Sprintf("%d", len(urls))+").\n")
		return
	}

	fmt.Fprintln(w, "Threads: "+threads+".\n") // Browser print

	// Retreive titles from pages
	titles, succeeded, failed := getTitle(urlCh, statCh, intThreads, urls)

	// Print the titles
	fmt.Fprintln(w, fmt.Sprintf("%d", succeeded)+" titles were found:\n")
	for i := 0; i < len(titles); i++ {
		fmt.Fprintln(w, titles[i])
	}
	fmt.Fprintln(w, "\nThe number of successful calls were: "+fmt.Sprintf("%d", succeeded)+".")
	fmt.Fprintln(w, "The number of failed calls were: "+fmt.Sprintf("%d", failed)+".")

	return
}

func getTitle(urlCh chan string, statCh chan string, threads int, urls []string) ([]string, int, int) {
	//URL and status channels, threads. Returns: titles, succesful calls and failed calls.

	quotient := len(urls) / threads
	remainder := len(urls) % threads
	var wg sync.WaitGroup

	for i := 0; i < quotient; i++ {
		wg.Add(threads)
		for j := 0; j < threads; j++ {
			log.Println("Fetching title from URL " + fmt.Sprintf("%d", threads*i+j))
			go parseHTML(urlCh, statCh, urls[threads*i+j], &wg)
		}
		wg.Wait()
	}
	wg.Wait()

	if remainder != 0 {
		wg.Add(remainder)
		for k := 0; k < remainder; k++ {
			log.Println("Fetching title from URL " + fmt.Sprintf("%d", threads*quotient+k))
			go parseHTML(urlCh, statCh, urls[threads*quotient+k], &wg)
		}
	}
	wg.Wait()

	// Get titles and status from the above calls to parseHTML from channels
	var titles []string
	failed := 0
	succeeded := 0

	for i := 0; i < len(urls); i++ {
		title := <-urlCh
		if title != "" {
			titles = append(titles, title)
		}
		status := <-statCh
		if status == "succeeded" {
			succeeded++
		}
		if status == "failed" {
			failed++
		}
	}

	fmt.Println("getTitle funcion exiting.")
	return titles, succeeded, failed
}

// Fetches website and finds title
func parseHTML(urlCh chan string, statCh chan string, URL string, wg *sync.WaitGroup) {
	fmt.Println("Executing parseHTML on " + URL)

	// Get the webpage--------------------------------

	resp, err := http.Get(URL)
	// Handle the error if there is one
	if err != nil {
		log.Println("Error fetching page " + URL)
		statCh <- "failed" //Pass back status for fail count
		urlCh <- ""
		wg.Done()
		return
	}

	statCh <- "succeeded" //Pass back status for fail count
	log.Println("Successfully fetched page " + URL)

	defer resp.Body.Close()
	// Reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error converting HTML to String for page " + URL)
		urlCh <- ""
		statCh <- "failed" //Pass back status for fail count
		wg.Done()
		return
	}

	// Store the HTML code as a string
	text := string(html)

	// Find the title ------------------------------------

	// RegEx for finding text between <title></title> tags
	re := regexp.MustCompile(`<title.*?>(.*)</title>`)

	submatchall := re.FindAllStringSubmatch(text, -1)
	for _, element := range submatchall {
		// Pass into channel ch
		urlCh <- element[1]
		fmt.Println(element[1])
	}

	wg.Done()
	fmt.Println("Finished Executing parseHTML on " + URL)
	return
}
