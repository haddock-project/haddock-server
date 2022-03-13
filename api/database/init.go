package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func Init() {
	var err error
	db, err = sql.Open("sqlite3", "data/sql.db")
	if err != nil {
		log.Fatalln("Unable to open a new database in the persistent data folder : \n", err)
	}
}
