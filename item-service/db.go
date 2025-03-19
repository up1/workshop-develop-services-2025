package api

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/XSAM/otelsql"
	_ "github.com/go-sql-driver/mysql"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
)

func ConnectDB() (*sql.DB, error) {
	// Read database connection details from environment variables
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	database := os.Getenv("DATABASE_NAME")

	// Connect to the MySQL database and tracing
	db, err := otelsql.Open("mysql", username+":"+password+"@tcp("+host+":3306)/"+database,
		otelsql.WithAttributes(semconv.DBSystemMySQL))
	if err != nil {
		return nil, err
	}

	// Register DB stats to meter
	err = otelsql.RegisterDBStatsMetrics(db, otelsql.WithAttributes(
		semconv.DBSystemMySQL,
	))
	if err != nil {
		log.Fatal(err)
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
