package main

import (
	"database/sql"
	"flag"
	"log"

	_ "github.com/lib/pq"
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

var db *sql.DB

func main() {
	var flagPort string
	flag.StringVar(&flagPort, "port", "8080", "Port for server setup.")
	var host string
	flag.StringVar(&host, "host", "dockerHost", "Host to connect to.")
	flag.Parse()

	dbConnect(host)

	server := newServer(flagPort)

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
	log.Printf("Server has shut down.")

	closeChannels()
	db.Close()
}
