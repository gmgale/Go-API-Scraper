package main

import (
	"fmt"
	"log"
	"sync"
)

// getTitle is a function that returns an array of titles from urls
// and an array of sucessful and failed calls.
// It will process in concurrent batches of "threads" number of Goroutines,
// then any remaining urls will be processed concurrently.
func getTitle(threads int) ([]string, [2]int) {

	// titles is a list of titles extracted from urls from function parseHTML.
	var titles []string

	// urlCh is for title data.
	urlCh := make(chan string)
	defer close(urlCh)

	// statCh is for GET url succ/fail counter.
	statCh := make(chan string)
	defer close(statCh)

	quotient := len(urls) / threads
	remainder := len(urls) % threads
	succeeded := 0
	failed := 0
	var wg sync.WaitGroup

	for i := 0; i < quotient; i++ {
		wg.Add(threads)
		for j := 0; j < threads; j++ {
			log.Println("Fetching title from URL " + fmt.Sprintf("%d", threads*i+j))
			go parseHTML(urlCh, statCh, urls[threads*i+j], &wg)
			select {
			case status := <-statCh:
				if status == statusSucceeded {
					succeeded++
				}
				if status == statusFailed {
					failed++
				}
			}
			title := <-urlCh
			if title != "" {
				titles = append(titles, title)
			}
		}
		wg.Wait()
	}
	wg.Wait()

	if remainder != 0 {
		wg.Add(remainder)
		for k := 0; k < remainder; k++ {
			log.Println("Fetching title from URL " + fmt.Sprintf("%d", threads*quotient+k))
			go parseHTML(urlCh, statCh, urls[threads*quotient+k], &wg)
			select {
			case status := <-statCh:
				if status == statusSucceeded {
					succeeded++
				}
				if status == statusFailed {
					failed++
				}
			}
			title := <-urlCh
			if title != "" {
				titles = append(titles, title)
			}
		}
	}
	wg.Wait()

	fmt.Println("getTitle funcion exiting.")

	var statusArr [2]int
	statusArr[0] = succeeded
	statusArr[1] = failed
	return titles, statusArr
}
