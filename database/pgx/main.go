package main

import (
	"context"
	_ "database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

const (
	host     = "127.0.0.1"
	port     = 5433
	user     = "yugabyte"
	password = "yugabyte"
	dbname   = "yugabyte"
)

func main() {

	yugaconnstring := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	var err error
	db, err = pgx.Connect(context.Background(), yugaconnstring)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close(context.Background())

	ctx := context.Background()

	var createStmt = `CREATE TABLE IF NOT EXISTS employees (id int PRIMARY KEY),
		name varchar,
		age int,
		language varchar)`
	if _, err := db.Exec(ctx, createStmt); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Created table employees")

	// Insert into the table.
	var insertStmt string = "INSERT INTO employees(name, age, language) VALUES ('John', 35, 'Go')"
	if _, err := db.Exec(ctx, insertStmt); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted data: %s\n", insertStmt)

	// Read from the table.
	var name string
	var age int
	var language string
	rows, err := db.Query(ctx, `SELECT name, age, language FROM employees WHERE id = 1`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	fmt.Printf("Query for id=1 returned: ")
	for rows.Next() {
		err := rows.Scan(&name, &age, &language)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Row[%s, %d, %s]\n", name, age, language)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
