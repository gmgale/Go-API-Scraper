package main

import (
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/renderer"
	"log"
	"net/http"
)

type rowStr struct {
	data string
	id   int
}

type localDataStr struct {
	data titleDataStr
	id   int
}

var rnd *renderer.Render

// dispResults is a function to pull data from the database
// and then print to the browser.
func dispResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	rows, _ := db.Query(`SELECT id, data FROM calls;`)

	// dbLocal will be a local pull of the database
	var dbLocal []localDataStr

	for rows.Next() {
		var row rowStr

		err := rows.Scan(&row.id, &row.data)
		if err != nil {
			log.Println("Error reading from db rows.")
		} else {
			var newTitleData titleDataStr

			byteData := []byte(row.data)
			if err := json.Unmarshal(byteData, &newTitleData); err != nil {
				log.Println("Error unmarshalling JSON.")
			}

			var x localDataStr
			x.id = row.id
			x.data = newTitleData
			dbLocal = append(dbLocal, x)
		}
	}

	for i := 0; i < len(dbLocal); i++ {
		fmt.Fprintln(w, "ID: ", dbLocal[i].id)
		fmt.Fprintln(w, dbLocal[i].data)
	}

	opts := renderer.Options{
		ParseGlobPattern: "./api/pages/*.html",
	}
	rnd = renderer.New(opts)

	rnd.HTML(w, http.StatusOK, "results", nil)

	rows.Close()
}
