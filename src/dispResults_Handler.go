package main

import (
	// "database/sql"
	//"encoding/json"

	"log"
	"net/http"
)

func dispResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	rows, _ := db.Query("SELECT data FROM calls LIMIT $1", 3)

	var data titleDataStr

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&data)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(data)
		//textData := json.Unmarshal(data []byte, titleDataStr)
	}
}
