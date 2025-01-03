// cmd/server/main.go
package main

import (
	"context"
	"net/http"
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
	cfg := config.LoadConfig()
	logger := logger.NewLogger()

	// Initialize components
	repo := repository.NewEventRepository(cfg.URL)
	eventUseCase := usecase.NewEventUseCase(repo)
	eventWorker := worker.NewEventWorker(100, eventUseCase)
	eventHandler := handlers.NewEventHandler(eventUseCase, eventWorker)

	// Start worker
	eventWorker.Start()

	// Setup server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: server.NewServer(eventHandler).Router(),
	}

	// Server run context
	serverCtx, serverStopCtx := context.WithCancel(context.Background())

	// Listen for syscall signals for process to interrupt/quit
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-sig

		// Shutdown signal with grace period of 30 seconds
		shutdownCtx, _ := context.WithTimeout(serverCtx, 30*time.Second)

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				logger.Fatal("graceful shutdown timed out.. forcing exit.", nil)
			}
		}()

		// Trigger graceful shutdown
		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			logger.Fatal("server shutdown error", err)
		}
		serverStopCtx()
	}()

	// Run the server
	logger.Info("Server is running on :8080")
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Fatal("server error", err)
	}

	// Wait for server context to be stopped
	<-serverCtx.Done()
}

// package main

// import (
// 	"net/http"

// 	"github.com/ratheeshkumar/event-processor/config"
// 	"github.com/ratheeshkumar/event-processor/logger"
// 	"github.com/ratheeshkumar/event-processor/pkg/handlers"
// 	"github.com/ratheeshkumar/event-processor/pkg/repository"
// 	"github.com/ratheeshkumar/event-processor/pkg/server"
// 	"github.com/ratheeshkumar/event-processor/pkg/usecase"
// 	"github.com/ratheeshkumar/event-processor/pkg/worker"
// )

// func main() {
// 	cfg := config.LoadConfig()
// 	logger := logger.NewLogger()

// 	// Initialize components
// 	repo := repository.NewEventRepository(cfg.URL)
// 	eventUseCase := usecase.NewEventUseCase(repo)
// 	eventWorker := worker.NewEventWorker(100, eventUseCase)
// 	eventHandler := handlers.NewEventHandler(eventUseCase, eventWorker)

// 	// Setup and start server
// 	srv := server.NewServer(eventHandler)
// 	logger.Info("Starting server on :8080")
// 	if err := http.ListenAndServe(":8080", srv.Router()); err != nil {
// 		logger.Fatal("Server failed to start", err)
// 	}
// }
