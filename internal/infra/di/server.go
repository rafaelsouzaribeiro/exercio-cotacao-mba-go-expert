package di

import (
	"database/sql"

	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/infra/database/sqlite/repository"
	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/usecase"
)

func NewServerDI(db *sql.DB) *usecase.UseCase {
	repository := repository.NewRepository(db)
	repository.Create()
	return usecase.NewUseCase(repository)
}
