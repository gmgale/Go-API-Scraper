package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	statusSucceeded = "succeeded"
	statusFailed    = "failed"

	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "webcalls"
)

// urls is a list of web addresses that will be processed.
var urls = [...]string{
	"https://www.result.si/projekti/",
	"https://www.result.si/o-nas/",
	"https://www.result.si/kariera/",
	"https://www.result.si/blog/",
}

var globalCallCounter = 0
var db *sql.DB

func main() {

	var flagPort string
	flag.StringVar(&flagPort, "flagPort", "8080", "Port for server setup. Default is 8080.")
	flag.Parse()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to PostgreSQL server!")

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
	log.Printf("DONE!")

	closeChannels()
	db.Close()
}
