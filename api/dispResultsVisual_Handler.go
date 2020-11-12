package main

import (

	"encoding/json"
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

func init() {
	opts := renderer.Options{
		ParseGlobPattern: "./web/*.html",
	}
	rnd = renderer.New(opts)
}

// dispResultsVisual is a function to pull data from the database
// and then print to the browser in formatted HTML.
func dispResultsVisual(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "html")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	rows, _ := db.Query(`SELECT id, data FROM calls;`)

	// dbLocal will be a local pull of the database
	var DBLocal []localDataStr

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
			DBLocal = append(DBLocal, x)
		}
	}
	rows.Close()
	log.Println("Data unmarshalled.")

	rnd.HTML(w, http.StatusOK, "results", DBLocal)
}
