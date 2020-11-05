package main

import (
	"encoding/json"

	_ "github.com/lib/pq"
)

func dbSend(data dataJSON) {
	newDataJSON, err := json.Marshal(data)

	sqlStatement := `
	INSERT INTO webcalls (data) 
	VALUES ($1)`
	_, err = db.Exec(sqlStatement, newDataJSON)
	if err != nil {
		panic(err)
	}
}
