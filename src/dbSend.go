package main

import (
	"encoding/json"
	_ "github.com/lib/pq"
)

func dbSend(newData titleDataStr) {
	newDataJSON, err := json.Marshal(newData)
	sqlStatement := `
	INSERT INTO calls (data) 
	VALUES ($1)`
	_, err = db.Exec(sqlStatement, newDataJSON)
	if err != nil {
		panic(err)
	}
}
