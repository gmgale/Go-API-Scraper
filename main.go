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
	"sync"

	//"strings"

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
	//log.Println("Channels for goroutines initilised.")

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

	// Split the title and print out
	/* for i := 0; i <= len(titles); i++ {
		x := strings.Split(titles[i], " ")
		fmt.Fprintln(w, x)
	} */
	fmt.Fprintln(w, titles)
	fmt.Fprintln(w, "The number of successful calls were: "+fmt.Sprintf("%d", succeeded)+".")
	fmt.Fprintln(w, "The number of failed calls were: "+fmt.Sprintf("%d", failed)+".")

	log.Println("getThreads func Exiting...")

}

func getTitle(urlCh chan string, statCh chan string, threads int) ([]string, int, int) {
	//URL and status channels, threads. Returns: titles, succesful calls and failed calls.
	url := 4
	quotient := url / threads
	remainder := url % threads

	var wgq sync.WaitGroup
	//wgq.Add(quotient)

	for i := 0; i < quotient; i++ {
		log.Println("i = " + fmt.Sprintf("%d", i))
		wgq.Add(threads)
		for j := 0; j < threads; j++ {
			log.Println("j = " + fmt.Sprintf("%d", j))
			log.Println("Fetching title from URL " + fmt.Sprintf("%d", threads*i+j))
			go parseHTML(urlCh, statCh, getURL(threads*i+j), &wgq)
		}
		wgq.Wait()
	}
	log.Println("Waiting wg q1")
	wgq.Wait()
	log.Println("Waiting wg q1 over")
	if remainder != 0 {
		log.Println("Remainder is " + fmt.Sprintf("%d", remainder))
		wgq.Add(remainder)
		for k := 0; k < remainder; k++ {
			log.Println("k = " + fmt.Sprintf("%d", threads*quotient+k))
			log.Println("Fetching title from URL " + fmt.Sprintf("%d", threads*quotient+k))
			go parseHTML(urlCh, statCh, getURL(threads*quotient+k), &wgq)
		}
	}
	log.Println("Waiting wg q2")
	wgq.Wait()
	log.Println("Waiting wg q2 over")

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
func parseHTML(urlCh chan string, statCh chan string, URL string, wgq *sync.WaitGroup) {
	fmt.Println("Executing parseHTML on " + URL)

	// Get the webpage--------------------------------

	resp, err := http.Get(URL)
	// Handle the error if there is one
	if err != nil {
		log.Println("Error fetching page " + URL)
		statCh <- "failed" //Pass back status for fail count
		panic(err)
	}
	statCh <- "succeeded" //Pass back status for fail count
	log.Println("Successfully fetched page " + URL)
	// Do this now so it won't be forgotten
	defer resp.Body.Close()
	// Reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error converting HTML to String for page " + URL)
		panic(err)
	}
	log.Println("Converted HTML to string for " + URL)

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
	wgq.Done()
	fmt.Println("Finished Executing parseHTML on " + URL)
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
