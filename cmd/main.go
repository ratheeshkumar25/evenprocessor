package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ratheeshkumar/event-processor/config"
	"github.com/ratheeshkumar/event-processor/logger"
	"github.com/ratheeshkumar/event-processor/pkg/handlers"
	"github.com/ratheeshkumar/event-processor/pkg/repository"
	"github.com/ratheeshkumar/event-processor/pkg/server"
	"github.com/ratheeshkumar/event-processor/pkg/usecase"
	"github.com/ratheeshkumar/event-processor/pkg/worker"
)

func main() {
	// Initialize components
	cfg := config.LoadConfig()
	logger := logger.NewLogger()

	// Setup application components
	repo := repository.NewEventRepository(cfg.URL)
	eventUseCase := usecase.NewEventUseCase(repo)
	eventWorker := worker.NewEventWorker(100, eventUseCase)
	eventHandler := handlers.NewEventHandler(eventUseCase, eventWorker)

	// Start worker
	eventWorker.Start()

	// Create and configure HTTP server
	httpServer := server.NewHTTPServer(
		logger,
		server.WithPort("8080"),
		// Add any middleware here if needed
		// server.WithMiddleware(middleware.Logger),
	)

	// Register routes
	httpServer.RegisterRoutes(eventHandler)

	// Setup graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Start server in a goroutine
	go func() {
		if err := httpServer.Start(); err != nil {
			logger.Fatal("Error starting server", err)
		}
	}()

	logger.Info("Server is running...")

	// Wait for interrupt signal
	<-stop

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Stop(ctx); err != nil {
		logger.Fatal("Error shutting down server", err)
	}

	// Stop worker
	eventWorker.Stop()

	logger.Info("Server stopped gracefully")
}
