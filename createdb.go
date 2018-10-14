package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "./data.db")
	sqlStmt := `
  create table ipflows (id integer not null primary key, sourceip text, destip text, protocol text, port text);
	delete from ipflows;
	`
	query, _ := db.Prepare(sqlStmt)
	query.Exec()
}
