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
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	app := fiber.FiberConn(&cfg.App)

	gormDb, err := gorm.DbConn(&cfg.Database)
	if err != nil {
		panic("can't connect to db")
	}

	app.Use(cors.New())

	go func() {
		slog.Info(fmt.Sprintf("Starting server on port %s", cfg.App.Port))
		if err := app.Listen(cfg.App.Port); err != nil {
			slog.Error("Failed to start server", slog.Any("error", err))
			stop()
		}
	}()

	server.Start(ctx, &cfg.App, app, gormDb) // add gormDb

	<-ctx.Done()
	slog.Info("Received shutdown signal, shutting down server...")

	if err := app.ShutdownWithContext(ctx); err != nil {
		slog.Error("Failed to gracefully shutdown server", slog.Any("error", err))
	} else {
		slog.Info("Server shutdown completed")
	}
}
