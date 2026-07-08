package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
)

func main() {
	var err error

	// open or create database
	db, err = sql.Open("sqlite3", "./users.db")

	if err != nil {
		log.Fatal(err)
	}

	db.Exec(`CREATE TABLE IF NOT EXISTS todos(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	task TEXT NOT NULL,
	done INTEGER DEFAULT 0
)`)

	port := os.Getenv("PORT")

	if port == "" {
		port = "1990"
	}
	http.HandleFunc("/{$}", homeHandler)
	http.HandleFunc("/newtask", taskHandler)

	log.Println("Starting Server on port:" + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
