package dal

import (
	"database/sql"
	"log"
)

var db *sql.DB

func Init(db *sql.DB, url string) {
	var err error
	db, err = sql.Open("sqlite3", url)
	if err != nil {
		log.Fatal("Error in initiating db connection QueryManager")
	}

}
