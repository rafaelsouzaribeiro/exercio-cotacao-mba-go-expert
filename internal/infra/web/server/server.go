package server

import (
	"net/http"

	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/infra/web/handlers"
	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/usecase"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}

func (s *Server) SetRoute(usecase *usecase.UseCase) {
	handler := handlers.NewHandler(usecase)
	err := handler.Create()

	if err != nil {
		panic(err)
	}
	s.mux.HandleFunc("/cotacao", handler.Cotacao)
}

func (s *Server) Start() error {
	return http.ListenAndServe(":8080", s.mux)
}
