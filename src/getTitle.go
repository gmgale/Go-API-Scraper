package main

import (
	"fmt"
	"log"
)

// getTitle is a function that returns an array of titles from urls
// and an array of sucessful and failed calls.
// It will process in concurrent batches of "threads" number of Goroutines,
// then any remaining urls will be processed concurrently.
func getTitle(threads int) titleDataStr {

	// titleData is a struct to hold an array of titles (string)
	// and a struct of status sucess/fail counts (int)
	titleData := titleDataStr{}

	quotient := len(urls) / threads
	remainder := len(urls) % threads

	for i := 0; i < quotient; i++ {
		wg.Add(threads)
		for j := 0; j < threads; j++ {
			log.Println("Fetching title from URL " + fmt.Sprintf("%d", threads*i+j))
			go parseHTML(urls[threads*i+j])
			select {
			case status := <-statCh:
				if status == statusSucceeded {
					titleData.status.success++
				}
				if status == statusFailed {
					titleData.status.fail++
				}
			}
			newTitle := <-urlCh
			if newTitle != "" {
				titleData.titles = append(titleData.titles, newTitle)
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
					titleData.status.success++
				}
				if status == statusFailed {
					titleData.status.fail++
				}
			}
			newTitle := <-urlCh
			if newTitle != "" {
				titleData.titles = append(titleData.titles, newTitle)
			}
		}
	}
	wg.Wait()

	fmt.Println("getTitle funcion exiting.")

	return titleData
}
