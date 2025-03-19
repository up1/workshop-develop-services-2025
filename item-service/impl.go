package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// ensure that we've conformed to the `ServerInterface` with a compile-time check
var _ ServerInterface = (*Server)(nil)

type Server struct {
	DB *sql.DB
}

func NewServer(db *sql.DB) Server {
	return Server{DB: db}
}

// GetItems implements ServerInterface.
func (s *Server) GetItems(ctx echo.Context) error {
	// Trace the request
	tracer := otel.GetTracerProvider()
	_ctx, span := tracer.Tracer("demo").Start(ctx.Request().Context(), "get-items-from-db")
	defer span.End()

	// Meter the request
	meter := otel.Meter("demo")
	commonAttrs := []attribute.KeyValue{
		attribute.String("key1", "value1"),
		attribute.String("key2", "value2"),
	}
	runCount, err := meter.Int64Counter("get_items_count", metric.WithDescription("The number of times the iteration ran"))
	if err != nil {
		log.Fatal(err)
	}
	runCount.Add(ctx.Request().Context(), 1, metric.WithAttributes(commonAttrs...))

	// Query the database for items
	rows, err := s.DB.QueryContext(_ctx, "SELECT id, name FROM items")
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to query items"})
	}
	defer rows.Close()
	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.Id, &item.Name); err != nil {
			return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to scan item"})
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Error iterating over items"})
	}
	// Return the items as JSON
	if len(items) == 0 {
		return ctx.JSON(http.StatusNotFound, map[string]string{"error": "No items found"})
	}
	// Return the items as JSON
	return ctx.JSON(http.StatusOK, items)
}
