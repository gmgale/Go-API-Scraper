package main

import (
	"fmt"
	"log"
	"net/http"
)

type idrowStr struct {
	data string
	id   int
}

// dispResults is a function to pull data from the database
// and then print to the browser.
func dispResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	rows, _ := db.Query(`SELECT id, data FROM calls;`)

	for rows.Next() {
		var row idrowStr

		err := rows.Scan(&row.id, &row.data)
		if err != nil {
			log.Println("Error reading from db rows.")
		} else {
			fmt.Fprintln(w, "Id: ", row.id)
			fmt.Fprintln(w, row.data)
		}
	}
	rows.Close()
}
