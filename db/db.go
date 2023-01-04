package db

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnDB() *sql.DB {
	dotenv := godotenv.Load()
	if dotenv != nil {
		panic("Error loading .env")
	}
	conn := os.Getenv("PGCONFIG")
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
