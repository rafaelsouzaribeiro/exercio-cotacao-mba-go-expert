package server

import (
	"net/http"

	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/infra/web/handlers"
	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/usecase/server"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) SetRoute(usecase *server.UseCaseServer) {
	handler := handlers.NewHandler(usecase)
	handler.Create()
	s.mux.HandleFunc("/cambio", handler.Cambio)
}

func (s *Server) Start() error {
	return http.ListenAndServe(":8080", s.mux)
}
