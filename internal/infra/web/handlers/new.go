package handlers

import (
	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/usecase/server"
)

type UsecaseServer struct {
	UsecaseServer *server.UseCaseServer
}

func NewHandler(usecase *server.UseCaseServer) *UsecaseServer {
	return &UsecaseServer{
		UsecaseServer: usecase,
	}
}
