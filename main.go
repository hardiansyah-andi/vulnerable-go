package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func main() {

	os.Remove("sqlite-database.db") // I delete the file to avoid duplicated records.
	// SQLite is a file based database.

	log.Println("Creating sqlite-database.db...")
	file, err := os.Create("sqlite-database.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("sqlite-database.db created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlite-database.db") // Open the created SQLite File
	defer sqliteDatabase.Close()                                     // Defer Closing the database
	createTable(sqliteDatabase)                                      // Create Database Tables

	insertUser(sqliteDatabase, "Liana Kim", "pass1")
	insertUser(sqliteDatabase, "Glen Rangel", "pass2")
	insertUser(sqliteDatabase, "Martin Martins", "secret")

	userId := os.Args[1]

	//executeVulnerableQuery(sqliteDatabase, userId)
	executeQuery(sqliteDatabase, userId)
}

func insertUser(db *sql.DB, name string, password string) {
	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO user(name, password) VALUES (?, ?)`
	statement, err := db.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(name, password)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func executeVulnerableQuery(db *sql.DB, userId string) error {
	sql := "SELECT * FROM user WHERE id = " + userId
	row, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var password string
		row.Scan(&id, &name, &password)
		log.Println("User: ", id, name, password)
	}
	return nil
}

func executeQuery(db *sql.DB, userId string) error {
	re := regexp.MustCompile("^[a-zA-Z0-9_]*$")
	if !re.MatchString(userId) {
		e := errors.New("Invalid input")
		log.Println(e)
		return e
	}
	sql := fmt.Sprintf("SELECT * FROM user WHERE id = %s", userId)
	row, err := db.Query(sql)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var name string
		var password string
		row.Scan(&id, &name, &password)
		log.Println("User: ", id, name, password)
	}

	fmt.Println(sql)
	return nil
}

func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE user (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"password" TEXT
	  );` // SQL Statement for Create Table

	log.Println("Create user table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("user table created")
}
