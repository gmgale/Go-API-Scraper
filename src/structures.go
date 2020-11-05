package main

import "net/http"

// titleDataStr is a structure to hold a list of titles (string)
// and a struct of returned success and failed counts (int)
type titleDataStr struct {
	titles []string
	status struct {
		success int
		fail    int
	}
}

type myServer struct {
	http.Server
	shutdownReq chan bool
	reqCount    uint32
}

type dataJSON struct {
	id      int
	time    string
	results []struct {
		url   string
		title string
	}
	threads  int
	succeded int
	failed   int
	duration string
}
