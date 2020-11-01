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

const (
	statusSucceeded = "succeeded"
	statusFailed    = "failed"
)

func main() {
	startServer()
}

// startServer sets up a localhost server using the gorilla/mux package
// and calls handlers for endpoints.
func startServer() {
	log.Println("startServer funtion called.")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api", topLevel)
	router.HandleFunc("/api/{Id=threads}", getThreads)
	log.Fatal(http.ListenAndServe(":8080", router))
}

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

	// urls is a list of web addresses that will be processed.
	urls := []string{
		"https://www.result.si/projekti/",
		"https://www.result.si/o-nas/",
		"https://www.result.si/kariera/",
		"https://www.result.si/blog/",
	}

	// urlChannel is for title data.
	urlCh := make(chan string, 4)
	defer close(urlCh)

	// statChannel is for GET url succ/fail count.
	statCh := make(chan string, 4)
	defer close(statCh)

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

	// Retreive titles and calls from pages ***

	titles, succeeded, failed := getTitle(urlCh, statCh, intThreads, urls)

	// Print the titles	***

	fmt.Fprintln(w, fmt.Sprintf("%d", succeeded)+" titles were found:\n")
	for i := 0; i < len(titles); i++ {
		fmt.Fprintln(w, titles[i])
	}
	fmt.Fprintln(w, "\nThe number of successful calls were: "+fmt.Sprintf("%d", succeeded)+".")
	fmt.Fprintln(w, "The number of failed calls were: "+fmt.Sprintf("%d", failed)+".")
	return
}

// getTitle is a function that returns titles from urls,
// sucessful and failed calls.
// It will process in concurrent batches of "threads" number of Goroutines,
// then any remaining urls will be processed concurrently.
func getTitle(urlCh chan string, statCh chan string, threads int, urls []string) ([]string, int, int) {

	//URL and status channels, threads. Returns: titles, succesful calls and failed calls. ***

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

	// titles is a list of titles extracted from urls from function parseHTML.
	var titles []string

	failed := 0
	succeeded := 0

	for i := 0; i < len(urls); i++ {
		title := <-urlCh
		if title != "" {
			titles = append(titles, title)
		}
		status := <-statCh
		if status == statusSucceeded {
			succeeded++
		}
		if status == statusFailed {
			failed++
		}
	}

	fmt.Println("getTitle funcion exiting.")
	return titles, succeeded, failed
}

// parseHTML is a function that extracts a title from a URL.
// It then uses go channels to send a success/fail varible and the title.
func parseHTML(urlCh chan string, statCh chan string, URL string, wg *sync.WaitGroup) {
	fmt.Println("Executing parseHTML on " + URL)

	resp, err := http.Get(URL)

	if err != nil {
		log.Println("Error fetching page " + URL)
		statCh <- statusFailed
		urlCh <- ""
		wg.Done()
		return
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error converting HTML to String for page " + URL)
		urlCh <- ""
		statCh <- statusFailed
		wg.Done()
		return
	}

	statCh <- statusSucceeded
	log.Println("Successfully fetched page " + URL)

	text := string(html)

	//re is the regular expression for finding the title in the string of HTML.
	re := regexp.MustCompile(`<title.*?>(.*)</title>`)

	submatchall := re.FindAllStringSubmatch(text, -1)
	for _, element := range submatchall {
		urlCh <- element[1]
		fmt.Println(element[1])
	}

	wg.Done()
	fmt.Println("Finished Executing parseHTML on " + URL)
	return
}
