package main

import "sync"

const (
	statusSucceeded = "succeeded"
	statusFailed    = "failed"
)

// urls is a list of web addresses that will be processed.
var urls = [...]string{
	"https://www.result.si/projekti/",
	"https://www.result.si/o-nas/",
	"https://www.result.si/kariera/",
	"https://www.result.si/blog/",
}

// urlCh is a channel for title data.
var urlCh = make(chan string)

// statCh is a channel for GET url succ/fail counter.
var statCh = make(chan string)

// wg is the WaitGroup used for batch processing Goroutines
// in getTitle and parseHTML functions
var wg sync.WaitGroup

func main() {

	defer close(urlCh)
	defer close(statCh)

	startServer()
}
