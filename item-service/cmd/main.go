package main

import (
	"log"

	"api"

	"github.com/labstack/echo/v4"
)

func main() {
	// Connect to the mysql database
	db, err := api.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal("Failed to close the database connection:", err)
		}
	}()

	// Create server
	server := api.NewServer(db)

	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	api.RegisterHandlers(e, &server)

	// And we serve HTTP until the world ends.
	log.Fatal(e.Start("0.0.0.0:8080"))
}
