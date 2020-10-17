package main

// Enter url of "localhost:8080/api" to begin
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"

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
	log.Fatal(http.ListenAndServe(":8080", router)) // This loops forever, but is OK... I think...
}

// Welcome displays splash screen at .../api
func topLevel(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!\n\nAppend'/x' to the URL (where x is a number 1-4), to enable concurrent threads/goroutines.")
}

// API Threads Endpoint ---------------------------------------------------------
func getThreads(w http.ResponseWriter, r *http.Request) {

	// Make channels for Go Routines
	// myChannel is return title data, statusChannel for GET url succ/fail count
	urlCh := make(chan string, 4)
	statCh := make(chan string, 4)
	defer close(urlCh)
	defer close(statCh)
	log.Println("Channels for goroutines initilised.")

	// Get threads from URL
	vars := mux.Vars(r)
	threads := vars["Id=threads"]
	fmt.Fprintln(w, "Threads: "+threads+".\n")      // Browser print
	log.Println("Threads set to: " + threads + ".") // Console print

	// Convert threads string to int
	intThreads, err := strconv.Atoi(threads)
	if err != nil {
		// Handle error
		fmt.Println(err)
		os.Exit(2)
	}

	// Retreive titles from pages
	titles, succeeded, failed := getTitle(urlCh, statCh, intThreads)

	fmt.Fprintln(w, titles)
	fmt.Fprintln(w, "The number of successful calls were: "+fmt.Sprintf("%d", succeeded)+".")
	fmt.Fprintln(w, "The number of failed calls were: "+fmt.Sprintf("%d", failed)+".")

}

func getTitle(urlCh chan string, statCh chan string, threads int) ([]string, int, int) {
	//url and status channel, threads. Returns titles, succesful calls and failed calls.

	// Go Routine concurrency logic goes here
	// Maybe change this to a loop rather than switch for worst case?
	switch threads {
	case 1:
		log.Println("Im case 1.")
		for i := 0; i <= 3; i++ {
			parseHTML(urlCh, statCh, getURL(i))
		}
	case 2:
		log.Println("Im case 2.")
		for i := 0; i < 2; i++ {
			go parseHTML(urlCh, statCh, getURL(2*i))
			go parseHTML(urlCh, statCh, getURL(2*i+1))
		}
	case 3:
		log.Println("Im case 3.")
		go parseHTML(urlCh, statCh, getURL(0))
		go parseHTML(urlCh, statCh, getURL(1))
		go parseHTML(urlCh, statCh, getURL(2))
		parseHTML(urlCh, statCh, getURL(3))
	case 4:
		log.Println("Im case 4.")
		go parseHTML(urlCh, statCh, getURL(0))
		go parseHTML(urlCh, statCh, getURL(1))
		go parseHTML(urlCh, statCh, getURL(2))
		go parseHTML(urlCh, statCh, getURL(3))
	default:
		log.Fatal("Thread input error. Out of bounds.")
	}

	// Get titles and status from the above calls to parseHTML from channels
	var titles []string
	failed := 0
	succeeded := 0

	for i := 0; i <= 3; i++ { // This will change to range of URLS length
		title := <-urlCh
		titles = append(titles, title)
		log.Println("Title " + fmt.Sprintf("%d", i+1) + " is: " + title) //SprintF converts to string
		status := <-statCh
		if status == "succeeded" {
			succeeded++
		}
		if status == "failed" {
			failed++
		}
	}

	log.Println(titles)

	log.Println("getTitle funcion exiting.")
	return titles, succeeded, failed
}

// Get website and finds titleng(
func parseHTML(urlCh chan string, statCh chan string, URL string) {
	fmt.Println("Executing parseHTML.")

	// Get the webpage--------------------------------

	resp, err := http.Get(URL)
	// Handle the error if there is one
	if err != nil {
		statCh <- "failed" //Pass back status for fail count
		panic(err)
	}
	statCh <- "succeeded" //Pass back status for fail count
	// Do this now so it won't be forgotten
	defer resp.Body.Close()
	// Reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// Store the HTML code as a string %s
	text := string(html)

	// Find the title ------------------------------------

	// RegEx for finding text between <title></title> tags
	re := regexp.MustCompile(`<title.*?>(.*)</title>`)

	submatchall := re.FindAllStringSubmatch(text, -1)
	for _, element := range submatchall {
		// Pass into channel ch
		urlCh <- element[1]
		//fmt.Println(element[1])
	}
	fmt.Println("Finished Executing parseHTML")
}

// Hold URLS in function
func getURL(index int) string {
	urls := [4]string{
		"https://www.result.si/projekti/",
		"https://www.result.si/o-nas/",
		"https://www.result.si/kariera/",
		"https://www.result.si/blog/"}
	return urls[index]
}
