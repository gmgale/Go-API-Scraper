package main

import (
	"fmt"
	"log"
)

// getTitle is a function that returns an array of titles from urls
// and an array of sucessful and failed calls.
// It will process in concurrent batches of "threads" number of Goroutines,
// then any remaining urls will be processed concurrently.
func getTitle(threads int) ([]string, [2]int) {

	// titles is a list of titles extracted from urls from function parseHTML.
	var titles []string

	quotient := len(urls) / threads
	remainder := len(urls) % threads
	succeeded := 0
	failed := 0

	for i := 0; i < quotient; i++ {
		wg.Add(threads)
		for j := 0; j < threads; j++ {
			log.Println("Fetching title from URL " + fmt.Sprintf("%d", threads*i+j))
			go parseHTML(urls[threads*i+j])
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
			go parseHTML(urls[threads*quotient+k])
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
