package main

import "sync"

// urlCh is a channel for title data.
var urlCh = make(chan string)

// statCh is a channel for GET url succ/fail counter.
var statCh = make(chan string)

// wg is the WaitGroup used for batch processing Goroutines
// in getTitle and parseHTML functions,
var wg sync.WaitGroup

// closeChannels is a functioin to closes down any open channels.
func closeChannels() {
	close(urlCh)
	close(statCh)
}
