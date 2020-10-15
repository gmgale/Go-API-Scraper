package main
 
import (
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/html"
	"github.com/gorilla/mux"
	
	//"io/ioutil"
)


func main() {

	startServer() // Getting stuck here ~~~~~~~~~~~~~~~~~~~~~~~~
	
}

//Set up a localhost server
func startServer(){
	log.Println("startServer funtion called.")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api", topLevel)
	router.HandleFunc("/api/{Id=threads}", getThreads)
	log.Fatal(http.ListenAndServe(":8080", router))
}

// Welcome displays splash screen
func topLevel(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!\n\nAppend'/x' to the URL (where x is a number 1-4), to enable concurrent threads/goroutines.")
}
// Get threads from URL
func getThreads(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	threads := vars["Id=threads"]
	fmt.Fprintln(w, "Threads: " + threads + ".\n")
	log.Println("Threads set to: " + threads + ".")

	//Channels for goroutines
	chUrls := make(chan string)
	chFinished := make(chan bool)
	log.Println("Channels for goroutines initilised.")

	//URLs to scrape
	urls := [4]string{
		"https://www.result.si/projekti/",
		"https://www.result.si/o-nas/",
		"https://www.result.si/kariera/",
		"https://www.result.si/blog/"}
	
	//
	foundTitles := make(map[string]bool)

	// Kick off the crawl process (concurrently)
	log.Println("Begin crawling.")

	for _, urls := range urls {
		go crawl(urls, chUrls, chFinished)
	}

	// Subscribe to both channels ----urls here should be titles?
	log.Println("Subscribing to channels.")
	for c := 0; c < len(urls); {
		select {
		case url := <-chUrls:
			foundTitles[url] = true
		case <-chFinished:
			c++
		}
	}
	
	// We're done! Print the results...
	fmt.Println("\nFound", len(foundTitles), "unique titles:")

	for url, _ := range foundTitles {
		fmt.Println(" - " + url)
	}

	close(chUrls)

}


// Helper function to pull the title attribute from a Token
func getTitle(t html.Token) (ok bool, title string) {
	// Iterate over token attributes until we find a "title"
	for _, a := range t.Attr {
		if a.Key == "title" {
			title = a.Val
			ok = true
		}
	}
	
	// "bare" return will return the variables (ok, title) as 
    // defined in the function definition
	return
}

// Extract all titles from a given webpage
func crawl(url string, ch chan string, chFinished chan bool) {
	resp, err := http.Get(url)
	

	defer func() {
		// Notify that we're done after this function
		log.Println("Crawl func chFinished")
		chFinished <- true
	}()

	if err != nil {
		fmt.Println("ERROR: Failed to find any titles:", url)
		return
	}

	b := resp.Body
	defer b.Close() // close Body when the function completes

	z := html.NewTokenizer(b)

	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()
			// Check if the token is an <title> tag
			isTitle := t.Data == "title"
			if !isTitle {
				continue
			}

			// Extract the title value, if there is one
			ok, title := getTitle(t)
			log.Println("Title found: " + title)
			ch <- title
			if !ok {
				continue
			}
		}
	}
}