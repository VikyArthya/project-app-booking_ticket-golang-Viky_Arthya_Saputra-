package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	connStr := "user=postgres password=superadmin dbname=ticketbooking sslmode=disable"
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatal("Failed to ping database: ", err)
	}
	log.Println("Database connected successfully")
}
