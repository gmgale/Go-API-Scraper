package main

import (
	"flag"
	"log"
)

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

func main() {

	var port string
	flag.StringVar(&port, "port", "8080", "Port for server setup. Default is 8080.")

	flag.Parse()

	log.Println("Port is set to: " + port)

	server := newServer(port)

	done := make(chan bool)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Printf("Listen and serve: %v", err)
		}
		done <- true
	}()

	server.waitShutdown()

	<-done
	log.Printf("DONE!")

	closeChannels()
}
