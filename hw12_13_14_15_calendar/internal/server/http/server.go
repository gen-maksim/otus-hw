package internalhttp

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server // TODO
}

type Logger interface { // TODO
}

type Application interface { // TODO
}

func NewServer(logger Logger, app Application) *Server {
	return &Server{}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {

		log.Printf("hello")

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Hi there!")
	}
}

func (s *Server) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:              ":8080",
		Handler:           loggingMiddleware(s),
		ReadHeaderTimeout: 90 * time.Second,
	}
	s.server = server

	server.ListenAndServe()

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.server.Shutdown(ctx)
	return nil
}
