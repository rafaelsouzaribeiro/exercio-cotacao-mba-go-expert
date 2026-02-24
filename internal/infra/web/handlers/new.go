package handlers

import (
	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/usecase"
)

type Usecase struct {
	UsecaseServer *usecase.UseCase
}

func NewHandler(usecase *usecase.UseCase) *Usecase {
	return &Usecase{
		UsecaseServer: usecase,
	}
}
