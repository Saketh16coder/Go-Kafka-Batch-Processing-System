package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
	connStr := "user=postgres password=postgres dbname=loadtesting port=5433 sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("DB Open Error:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("DB Connection Error:", err)
	}

	DB.SetMaxOpenConns(25)
	DB.SetMaxIdleConns(10)

	log.Println(" Connected to PostgreSQL")
}
