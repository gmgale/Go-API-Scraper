package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sync"
)

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
