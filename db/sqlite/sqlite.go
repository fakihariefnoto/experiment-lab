package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func main() {
	fileName := "user.db"
	createDB(fileName)
	connect(fileName)

}

func createDB(dbName string) {
	if _, err := os.Stat(dbName); os.IsExist(err) {
		os.Remove(dbName)
	}
	os.Create(dbName)
}

func connect(dbName string) (err error) {
	db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func disconnect() {
	if db != nil {
		db.Close()
	}
}

func createTable() {

}

func altertable() {

}

func insertQuery() {

}

func updateQuery() {

}

func upsertQuery() {

}

func deleteQuery() {

}

func selectQuery() {

}
