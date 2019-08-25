package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"

	"github.com/bpmericle/go-webservice/internal/handlers"
	"github.com/bpmericle/go-webservice/internal/logger"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

// Server represents the web server hosting the service
type Server struct {
	Port int
}

// ListenAndServe will start the web server and listen for requests
func (s *Server) ListenAndServe() error {

	// setup CHI router
	r := chi.NewRouter()

	// setup middlewares
	r.Use(middleware.Heartbeat("/ping")) // allows LB to verify service up
	r.Use(middleware.RequestID)          // ensures a request ID is logged
	r.Use(logger.NewStructuredLogger())  // uses structured logging like our app (logs only at debug level)
	r.Use(middleware.Recoverer)          // handles any unhandles errors and returns a 500

	// setup supported routes
	r.Get("/health", handlers.Health)

	address := fmt.Sprintf(":%d", s.Port)
	log.WithField("address", address).Info("starting server")

	return http.ListenAndServe(address, r)
}
