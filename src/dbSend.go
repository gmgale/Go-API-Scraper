package main

import (
	"encoding/json"
	_ "github.com/lib/pq"
	"log"
)

func dbSend(newData titleDataStr) {

	log.Println(newData)
	var newDataJSON []byte
	newDataJSON, err := json.MarshalIndent(newData, "", "	")

	sqlStatement := `
	INSERT INTO calls (data) 
	VALUES ($1);`
	_, err = db.Exec(sqlStatement, newDataJSON)
	if err != nil {
		panic(err)
	}
}
