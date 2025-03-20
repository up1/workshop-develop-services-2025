package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"

	"api"

	"github.com/labstack/echo/v4"
	slogotel "github.com/remychantenay/slog-otel"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"
	"go.opentelemetry.io/otel/sdk/resource"
)

func main() {
	// Initialize OpenTelemetry
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	conn, err := api.InitConn()
	if err != nil {
		log.Fatal(err)
	}

	res, err := resource.New(ctx,
		resource.WithAttributes(
			// The service name used to display traces in backends
			api.ServiceName,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	shutdownTracerProvider, err := api.InitTracerProvider(ctx, res, conn)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdownTracerProvider(ctx); err != nil {
			log.Fatalf("failed to shutdown TracerProvider: %s", err)
		}
	}()

	// Initialize the meter provider
	shutdownMeterProvider, err := api.InitMeterProvider(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdownMeterProvider(ctx); err != nil {
			log.Fatalf("failed to shutdown MeterProvider: %s", err)
		}
	}()

	// Initialize the log provider
	shutdownLogProvider, err := api.InitLogProvider(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := shutdownLogProvider(ctx); err != nil {
			log.Fatalf("failed to shutdown LogProvider: %s", err)
		}
	}()

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

	// Configure slog with OpenTelemetry handler
	slog.SetDefault(slog.New(slogotel.OtelHandler{
		Next: slog.NewJSONHandler(os.Stdout, nil),
	}))
	logger := slog.Default()
	logger = logger.With("service", "item-service")

	// Create server
	server := api.NewServer(db)

	e := echo.New()
	// Middleware for OpenTelemetry
	e.Use(otelecho.Middleware("item-service"))

	e.GET("/health", func(c echo.Context) error {
		logger.InfoContext(c.Request().Context(), "/health endpoint called", "locale", "th_TH")
		return c.String(200, "OK")
	})

	api.RegisterHandlers(e, &server)

	// And we serve HTTP until the world ends.
	log.Fatal(e.Start("0.0.0.0:8080"))
}
