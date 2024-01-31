package items

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func init() {
	var db *sql.DB
	var err error
	db, err = sql.Open("sqlite3", "./sqlite3db.db")
	if err != nil {
		log.Fatal("error opening ./sqlite3db.db")
	}
	statement, _ := db.Prepare(`
	create table if not exists  todo(
		"id" integer not null primary key autoincrement,
		"title" text,
		"description" text
	)	
	`)
	statement.Exec()
}
