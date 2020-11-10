package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

//dbSend is a function to send data to the database.
func dbSend(newData titleDataStr) {

	var newDataJSON []byte
	newDataJSON, err := json.MarshalIndent(newData, "", "	")

	sqlStatement := `
	INSERT INTO calls (data) 
	VALUES ($1);`
	_, err = db.Exec(sqlStatement, newDataJSON)
	if err != nil {
		fmt.Printf("Error sending new data to database.\n%v", err)
	}
}
