package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

type Server struct{}

func NewServer() Server {
	return Server{}
}

// GetItems implements ServerInterface.
func (s *Server) GetItems(ctx echo.Context) error {
	// Simulate fetching items from a database or other source
	items := []Item{
		{Id: 1, Name: "Item 1"},
		{Id: 2, Name: "Item 2"},
	}

	return ctx.JSON(http.StatusOK, items)
}
