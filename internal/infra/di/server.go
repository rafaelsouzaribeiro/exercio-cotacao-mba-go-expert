package di

import (
	"database/sql"

	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/infra/database/sqlite/repository"
	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/usecase/server"
)

func NewServerDI(db *sql.DB) *server.UseCaseServer {
	repository := repository.NewRepository(db)
	return server.NewServerUsecase(repository)
}
