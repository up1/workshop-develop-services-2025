package api

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	// Read database connection details from environment variables
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	database := os.Getenv("DATABASE_NAME")
	// Connect to the MySQL database
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":3306)/"+database)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	// Check if the connection is successful
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
