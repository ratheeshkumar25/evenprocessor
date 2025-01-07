package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/ratheeshkumar/event-processor/logger"
	"github.com/ratheeshkumar/event-processor/pkg/handlers"
)

type HTTPServer struct {
	router     *mux.Router
	server     *http.Server
	logger     *logger.Logger
	port       string
	middleware []mux.MiddlewareFunc
}

type HTTPServerOption func(*HTTPServer)

// WithPort sets the server port
func WithPort(port string) HTTPServerOption {
	return func(s *HTTPServer) {
		s.port = port
	}
}

// WithMiddleware adds middleware to the server
func WithMiddleware(middleware ...mux.MiddlewareFunc) HTTPServerOption {
	return func(s *HTTPServer) {
		s.middleware = append(s.middleware, middleware...)
	}
}

// NewHTTPServer creates a new HTTP server instance
func NewHTTPServer(logger *logger.Logger, options ...HTTPServerOption) *HTTPServer {
	server := &HTTPServer{
		router: mux.NewRouter(),
		logger: logger,
		port:   "8080",
	}

	// Apply options
	for _, option := range options {
		option(server)
	}

	// Apply middleware
	for _, m := range server.middleware {
		server.router.Use(m)
	}

	server.server = &http.Server{
		Addr:         fmt.Sprintf(":%s", server.port),
		Handler:      server.router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return server
}

// RegisterRoutes registers the API routes
func (s *HTTPServer) RegisterRoutes(handler *handlers.EventHandler) {
	s.router.HandleFunc("/api/events", handler.HandleEvent).Methods(http.MethodPost)
	// Add more routes as needed
}

// Start starts the HTTP server
func (s *HTTPServer) Start() error {
	s.logger.Info(fmt.Sprintf("Starting server on port %s", s.port))
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}

// Stop gracefully shuts down the server
func (s *HTTPServer) Stop(ctx context.Context) error {
	s.logger.Info("Shutting down server...")
	return s.server.Shutdown(ctx)
}
