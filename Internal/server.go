package internal

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

type CustomHandler struct {
}

func (s *Server) Run(port string, handler http.Handler) error {
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		http.Get(s.httpServer)
	})

	http.HandleFunc("/Shutdown", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		cancel()
		s.httpServer.Shutdown(ctx)
	})

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   20 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}
