package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "db"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "webcalls"
)

// dbConnect is a function to connect to the database.
func dbConnect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			fmt.Sprintln("Connecting to database. Tries remaining: ", i)
		}
		time.Sleep(5)
	}
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to PostgreSQL server!")
}
