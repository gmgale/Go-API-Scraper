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
}

func dispResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	rows, _ := db.Query(`SELECT data FROM calls;`)

	for rows.Next() {
		var row rowStr

		err := rows.Scan(&row.data)
		if err != nil {
			log.Println("ERROR 1")
		} else {
			fmt.Fprintln(w, row)
			log.Println(row)
		}
	}

	defer rows.Close()

}
