package main

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type jsondata struct {
	File string
	Hash string
}

func dbquery() {
	db, _ := sql.Open("sqlite3", "./data.db")
	defer db.Close()
	rows, err := db.Query("select id, sourceip, destip, protocol, port from ipflows")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var sourceip string
		var destip string
		var protocol string
		var port string
		err = rows.Scan(&id, &sourceip, &destip, &protocol, &port)
		if err != nil {
			log.Fatal(err)
		}
		// need to decide how the netflow data is going to be handled and stored.
		// for now lets print out the data.
		fmt.Println(id,sourceip,destip,protocol,port)
	}
}

func main() {
	dbquery()
}
