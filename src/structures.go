package main

import "net/http"

type myServer struct {
	http.Server
	shutdownReq chan bool
	reqCount    uint32
}

type titleDataStr struct {
	Time     string
	Threads  int
	Results  []urlTitleStr
	Status   statusStr
	Duration string
}

type urlTitleStr struct {
	Url   string
	Title string
}

type statusStr struct {
	Succeeded int
	Failed    int
}
