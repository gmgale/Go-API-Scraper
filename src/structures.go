package main

// titleDataStr is a structure to hold a list of titles (string)
// and a struct of returned success and failed counts (int)
type titleDataStr struct {
	titles []string
	status struct {
		success int
		fail    int
	}
}
