package server

import (
	"github.com/gorilla/mux"
	"github.com/ratheeshkumar/event-processor/pkg/handlers"
)

type Server struct {
	router  *mux.Router
	handler *handlers.EventHandler
}

func NewServer(handler *handlers.EventHandler) *Server {
	s := &Server{
		router:  mux.NewRouter(),
		handler: handler,
	}
	s.setupRoutes()
	return s
}

func (s *Server) setupRoutes() {
	s.router.HandleFunc("/api/events", s.handler.HandleEvent).Methods("POST")
}

func (s *Server) Router() *mux.Router {
	return s.router
}
