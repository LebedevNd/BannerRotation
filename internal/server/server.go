package internalhttp

import (
	"context"
	"fmt"
	"net/http"
)

type Server struct {
	server *http.Server
}

type MyHandler struct{}

func NewServer(host string, port int) *Server {
	handler := &MyHandler{}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handler.testHandler)
	fmt.Println("server will start on " + fmt.Sprintf("%s:%d", host, port))

	logger := NewLogger(mux)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: logger,
	}

	return &Server{
		server,
	}
}

func (s *Server) Start(ctx context.Context) error {
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}

	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.server.Close()
	return err
}

func (s *MyHandler) testHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello Rotator!"))
	if err != nil {
		return
	}
}
