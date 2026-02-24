package usecase

import "github.com/rafaelsouzaribeiro/exercicio-cotacao-mba-go-expert/internal/repository"

type UseCase struct {
	IRepository repository.IRepository
}

func NewUseCase(repo repository.IRepository) *UseCase {
	return &UseCase{
		IRepository: repo,
	}
}
