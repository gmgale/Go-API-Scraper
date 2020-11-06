package main

import (
	"encoding/json"

	_ "github.com/lib/pq"
)

func dbSend(newData titleDataStr) {
	titleDataStr, err := json.Marshal(newData)

	sqlStatement := `
	INSERT INTO webcalls (data) 
	VALUES ($1)`
	_, err = db.Exec(sqlStatement, titleDataStr)
	if err != nil {
		panic(err)
	}
}
