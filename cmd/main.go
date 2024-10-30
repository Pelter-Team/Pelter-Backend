package main

import (
	"Pelter_backend/internal/config"
	"Pelter_backend/internal/pkg/fiber"
	"Pelter_backend/internal/pkg/gorm"
	"Pelter_backend/internal/server"
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	cfg := config.LoadConfig()
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM) // listens for interrupt signals
	defer stop()                                                                                           // gracefully shutdown

	app := fiber.FiberConn(&cfg.App) // fiber app init

	gormDb, err := gorm.DbConn(&cfg.Database) // gorm db init
	if err != nil {
		panic("can't connect to db")
	}

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins:     cfg.App.Origin,
	}))

	go func() { // start server
		slog.Info(fmt.Sprintf("Starting server on port %s", cfg.App.Port))
		if err := app.Listen(cfg.App.Port); err != nil {
			slog.Error("Failed to start server", slog.Any("error", err))
			stop()
		}
	}()

	server.Start(ctx, &cfg.App, app, gormDb)

	<-ctx.Done() // wait for interrupt signal
	slog.Info("Received shutdown signal, shutting down server...")

	if err := app.ShutdownWithContext(ctx); err != nil {
		slog.Error("Failed to gracefully shutdown server", slog.Any("error", err))
	} else {
		slog.Info("Server shutdown completed")
	}
}
