package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {
	dsn := "root:parkhyunseok123@tcp(127.0.0.1:3306)/bookInfo"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Ping failed:", err)
	}

	return db
}
