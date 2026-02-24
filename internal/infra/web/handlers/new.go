package handlers

import (
	"github.com/rafaelsouzaribeiro/exercicio-cotacao-mba-go-expert/internal/usecase"
)

type Usecase struct {
	Usecase *usecase.UseCase
}

func NewHandler(usecase *usecase.UseCase) *Usecase {
	return &Usecase{
		Usecase: usecase,
	}
}
