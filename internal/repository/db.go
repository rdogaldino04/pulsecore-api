package repository

import (
	"database/sql"
	"fmt"
	"log"
	"pulsecore-api/internal/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.GetEnv("DB_HOST"),
		config.GetEnv("DB_PORT"),
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASSWORD"),
		config.GetEnv("DB_NAME"),
		config.GetEnv("DB_SSLMODE"),
	)

	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal(err)
	}
}
