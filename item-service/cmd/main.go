package main

import (
	"log"

	"api"

	"github.com/labstack/echo/v4"
)

func main() {
	server := api.NewServer()

	e := echo.New()

	e.GET("/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	api.RegisterHandlers(e, &server)

	// And we serve HTTP until the world ends.
	log.Fatal(e.Start("0.0.0.0:8080"))
}
