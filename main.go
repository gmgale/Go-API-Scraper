package main
 
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
	"os"
)

func main() {

	//Make channels for Go Routines
	myChannel := make(chan string)
	defer close(myChannel)
	/* log.Println("Channels for goroutines initilised.") */
	startServer() 
}

//Set up a localhost server
func startServer(){
	log.Println("startServer funtion called.")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api", topLevel)
	router.HandleFunc("/api/{Id=threads}", getThreads)
	log.Fatal(http.ListenAndServe(":8080", router))			//This loops forever, but is OK!
}
// Welcome displays splash screen
func topLevel(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!\n\nAppend'/x' to the URL (where x is a number 1-4), to enable concurrent threads/goroutines.")
}


// Get threads from URL
func getThreads(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	threads := vars["Id=threads"]
	fmt.Fprintln(w, "Threads: " + threads + ".\n")			// screen print
	log.Println("Threads set to: " + threads + ".")			// console print

	//Convert string to int
	intThreads, err := strconv.Atoi(threads)
    if err != nil {
        // handle error
        fmt.Println(err)
        os.Exit(2)
	}	
	
	//Retreive titles from pages
	titles := getTitle(intThreads)
	
	//Print the results
	log.Println("The titles are:")
	log.Println(titles)
}

func getURL(index int) string{
	urls := [4]string{
		"https://www.result.si/projekti/",
		"https://www.result.si/o-nas/",
		"https://www.result.si/kariera/",
		"https://www.result.si/blog/"}
	return urls[index]
}

func getTitle(threads int) []string{
	titles := []string{}

	// Go Routine concurrency logic goes her --------		:)
	// Maybe change this to a loop rather than switch?
	switch threads {
	case 1:
		fmt.Println("Im case 1.")
		for i := 0; i<=4; i++{
			go parseHTML(getURL(i))
		}
	case 2:
		fmt.Println("Im case 2.")
		for i := 0; i<=4; i=i+2{
			go parseHTML(getURL(i))
			go parseHTML(getURL(i+1))
		}
	case 3:
		fmt.Println("Im case 3.")
		for i := 0; i<=4; i=i+2{
			go parseHTML(getURL(i))
			go parseHTML(getURL(i+1))
			go parseHTML(getURL(i+2))
		}
			go parseHTML(getURL(3))
	case 4:
		fmt.Println("Im case 4.")
		go parseHTML(getURL(0))
		go parseHTML(getURL(1))
		go parseHTML(getURL(2))
		go parseHTML(getURL(3))
	}
	
	return titles
}

func parseHTML(URL string)string{
	//Use some ReGex ro find <title> tag contense
	title := ""
	
	return title
}