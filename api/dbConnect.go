package main

import (
	"database/sql"
	"fmt"
	"os"
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
	for i := 5; i >= 0; i-- {
		db, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			fmt.Printf("Error connecting to database %v\n Tries remaining: %d", err, i)
			time.Sleep(5)
		}
		if err == nil {
			break
		}
	}

	for i := 5; i >= 0; i-- {
		err = db.Ping()
		if err != nil {
			fmt.Printf("Error pinging database %v. Tries remaining: %d\n", err, i)
			time.Sleep(5)
		}
		if err == nil {
			break
		}
		if i == 0 {
			fmt.Printf("Error pining database. Use 'localhost' for testing and 'db' for docker.\n %v", err)
			os.Exit(1)
		}
	}
	fmt.Println("Successfully connected to PostgreSQL server!")
}
