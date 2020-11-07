package main

import (
	// "database/sql"
	//"encoding/json"

	"fmt"
	"log"
	"net/http"
)

type rowStr struct {
	data string
	id   int
}

func dispResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	rows, _ := db.Query(`SELECT id, data FROM calls;`)

	for rows.Next() {
		var row rowStr

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
