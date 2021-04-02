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

type (
	// User to storing user table data
	User struct {
		ID     int    `json:"user_id"`
		Name   string `json:"full_name"`
		Sex    string `json:"sex"`
		Age    int    `json:"age"`
		Hobby  string `json:"hobby"`
		Occupt string `json:"occupations"`
	}
)

func main() {
	fileName := "user.db"
	createDB(fileName)
	connect(fileName)
	createTable()
	alterTable()
	insertQuery()
	selectQuery()
	updateQuery()
	selectQuery()
	upsertQuery()
	selectQuery()
	updateQuery2()
	selectQuery()
	deleteQuery()
	selectQuery()
	disconnect()

}

func createDB(dbName string) {
	log.Println("createDB ", dbName)
	if _, err := os.Stat(dbName); os.IsExist(err) {
		os.Remove(dbName)
	}
	os.Create(dbName)
}

func connect(dbName string) (err error) {
	log.Println("connect ", dbName)
	db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal("connect, ", err)
	}
	return
}

func disconnect() {
	log.Println("disconnect")
	if db != nil {
		db.Close()
	}
}

func createTable() {
	log.Println("createTable user")
	syntax := `
	CREATE TABLE IF NOT EXISTS user (
		user_id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		sex TEXT NOT NULL,
		age TEXT NOT NULL
	);
	`
	// phone TEXT NOT NULL UNIQUE
	_, err := db.Exec(syntax)
	if err != nil {
		log.Fatalln("createTable, ", err)
	}
}

func alterTable() {
	log.Println("altertable user to users")
	syntax := `
		ALTER TABLE user RENAME TO users;
		ALTER TABLE users RENAME COLUMN name TO full_name;
		ALTER TABLE users ADD COLUMN hobby TEXT NULL;
		ALTER TABLE users ADD COLUMN occupations TEXT NULL;
	`
	_, err := db.Exec(syntax)
	if err != nil {
		log.Fatalln("altertable, ", err)
	}
}

/*

Fakih | Male   | 17 | Swimming | Swimmer
Hikaf | Female | 16 | - | Office Assistant
Kafhi | Male   | 11 | Jogging | -
Fhika | Female | 20 | - | -
Ikfha | Female | 55 | - | -

*/

func insertQuery() {
	log.Println("insertQuery")
	syntax := `
		INSERT INTO users(full_name,sex,age,hobby,occupations) 
			VALUES
				("Fakih", "Male", 17, "Swimming", "Swimmer"),
				("Hikaf", "Female", 16, "-", "Office Assistant"),
				("Kafhi", "Male", 11, "Jogging", ""),
				("Fhika", "Female", 20, "-", "-"),
				("Ikfha", "Female", 55, "-", "-");
	`
	_, err := db.Exec(syntax)
	if err != nil {
		log.Fatalln("insertQuery, ", err)
	}
}

func updateQuery() {
	log.Println("updateQuery")
	syntax := `
		UPDATE
			users
		SET
			full_name = "Fakih Coy"
		WHERE
			full_name = "Fakih"
	`
	_, err := db.Exec(syntax)
	if err != nil {
		log.Fatalln("updateQuery, ", err)
	}
}

/*

Fakih 1 | Male   | 17 | Swimming | Swimmer
Hikaf 2 | Female | 16 | - | Office Assistant
Kafhi 3 | Male   | 11 | Jogging | -
Fhika 4 | Female | 20 | - | -
Ikfha 5 | Female | 55 | - | -

*/

func upsertQuery() {
	log.Println("upsertQuery")
	syntax := `

	INSERT OR REPLACE INTO users (user_id, full_name, hobby, occupations, sex, age)
		VALUES (1, 'Fakih 1', 'Swimming', 'Swimmer Pro', 'Male', 13);

	INSERT INTO users(full_name,sex,age,hobby,occupations)
		VALUES
			("Hikaf 2", "Female", 16, "-", "Office Secretary"),
			("Kafhi 3", "Male", 11, "Basketball", ""),
			("Fhika 4 ", "Female", 20, "-", "-"),
			("Ikfha 5", "Female", 55, "-", "-");
	`
	_, err := db.Exec(syntax)
	if err != nil {
		log.Fatalln("upsertQuery, ", err)
	}
}

func updateQuery2() {
	log.Println("updateQuery2")
	syntax := `
		UPDATE
			users
		SET
			hobby = "Carambole"
		WHERE
			lower(full_name) LIKE '%ikfha%'
	`
	_, err := db.Exec(syntax)
	if err != nil {
		log.Fatalln("updateQuery2, ", err)
	}
}

/*
	Delete Ikfha
	Delete user with age 11
*/

func deleteQuery() {
	log.Println("deleteQuery Ikfha")
	syntax := `
		DELETE FROM
			users
		WHERE
			full_name = "Ikfha"
	`
	_, err := db.Exec(syntax)
	if err != nil {
		log.Fatalln("deleteQuery Ikfha, ", err)
	}

	log.Println("deleteQuery user with age 11")
	syntax = `
		DELETE FROM
			users
		WHERE
			age = 11
	`
	_, err = db.Exec(syntax)
	if err != nil {
		log.Fatalln("deleteQuery user with age 11, ", err)
	}
}

func selectQuery() {
	log.Println("selectQuery")
	syntaxBelow18 := `
		SELECT 
			user_id,
			full_name,
			sex,
			age,
			hobby,
			occupations
		FROM 
			users
		WHERE 
			age < 18

	`
	users := make([]User, 0)
	rows, err := db.Query(syntaxBelow18)
	if err != nil {
		log.Fatalln("selectQuery syntaxBelow18, ", err)
	}
	defer rows.Close()
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.ID, &u.Name, &u.Sex, &u.Age, &u.Hobby, &u.Occupt)
		if err != nil {
			log.Fatal("selectQuery syntaxBelow18, ", err)
		}
		users = append(users, u)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal("selectQuery syntaxBelow18, ", err)
	}
	if len(users) > 0 {
		log.Printf("User with age below 18 total %v \n", len(users))
		for _, u := range users {
			log.Printf("%#v\n", u)
		}
	}

	syntaxAbove18 := `
	SELECT 
		user_id,
		full_name,
		sex,
		age,
		hobby,
		occupations
	FROM 
		users
	WHERE 
		age > 18

	`
	users = make([]User, 0)
	rows, err = db.Query(syntaxAbove18)
	if err != nil {
		log.Fatalln("selectQuery syntaxAbove18, ", err)
	}
	defer rows.Close()
	for rows.Next() {
		u := User{}
		err = rows.Scan(&u.ID, &u.Name, &u.Sex, &u.Age, &u.Hobby, &u.Occupt)
		if err != nil {
			log.Fatal("selectQuery syntaxAbove18, ", err)
		}
		users = append(users, u)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal("selectQuery syntaxAbove18, ", err)
	}
	if len(users) > 0 {
		log.Printf("User with age above 18 total %v \n", len(users))
		for _, u := range users {
			log.Printf("%#v\n", u)
		}
	}

}
