package server

import (
	"log"
	"net/http"
	"time"

	"morze_converter/internal/handlers"
)

type Server struct {
	logger *log.Logger
	http   *http.Server
}

func NewServer(logger *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/upload", handlers.UploadHandler)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger: logger,
		http:   srv,
	}
}

func (s *Server) Start() error {
	return s.http.ListenAndServe()
}
