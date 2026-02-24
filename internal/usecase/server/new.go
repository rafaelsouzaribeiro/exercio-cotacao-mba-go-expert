package server

import "github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/repository"

type UseCaseServer struct {
	IRepository repository.IServerRepository
}

func NewServerUsecase(repo repository.IServerRepository) *UseCaseServer {
	return &UseCaseServer{
		IRepository: repo,
	}
}
