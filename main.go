package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var db *sql.DB
	var err error
	_, err = os.Stat("./sqlite3db.db")
	if err != nil {
		os.Create("sqlite3db.db")
	}
	db, err = sql.Open("sqlite3", "./sqlite3db.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	createTodo(db)
	var statement *sql.Stmt
	statement, err = db.Prepare("insert into todo(title,description) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = statement.Exec("titleasdf", "descripsssstion")
	if err != nil {
		log.Fatal(err)
	}
	row, err := db.Query("select title,description from todo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(row)
	for row.Next() {
		var title string
		var description string
		row.Scan(&title, &description)
		log.Println(title, description)
	}
}
func createTodo(db *sql.DB) {
	statement, _ := db.Prepare(`
	create table if not exists  todo(
		"id" integer not null primary key autoincrement,
		"title" text,
		"description" text
	)	
	`)
	statement.Exec()
	log.Println("todo created")
}
