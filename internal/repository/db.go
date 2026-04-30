package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "pulsecore.db")
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

}

func Migrate() {
	query := `
	CREATE TABLE IF NOT EXISTS pessoas (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		nome TEXT NOT NULL
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
