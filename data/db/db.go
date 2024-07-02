package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const dbName = "database.db"

func getDB() *sql.DB {
	if _, err := os.Stat(dbName); os.IsNotExist(err) {
		fmt.Println("Creating database...")
		createDatabase(dbName)
	} else if err != nil {
		log.Fatalf("Failed to check if database exists: %v", err)
	} else {
		fmt.Println("Database already exists.")
	}

	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	// defer closeDB(db)
	fmt.Println("Database opened successfully.")

	return db
}

// func closeDB(db *sql.DB) {
// 	defer db.Close()
// 	fmt.Println("Database closed successfully.")
// }

func createDatabase(dbName string) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
	 log.Fatalf("Failed to create database: %v", err)
	}
	defer db.Close()
 
	createTableSQL := `CREATE TABLE IF NOT EXISTS groups (
	 "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	 "identifier" integer
	);`
 
	_, err = db.Exec(createTableSQL)
	if err != nil {
	 log.Fatalf("Failed to create table: %v", err)
	}
 
	fmt.Println("Database and table created successfully.")
 }

var DB = getDB()