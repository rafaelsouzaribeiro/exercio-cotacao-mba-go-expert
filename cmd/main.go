package main

import (
	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/infra/database/sqlite/connection"
	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/infra/database/sqlite/repository"
	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/infra/web/server"
	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/usecase"
)

func main() {
	server := server.NewServer()
	db, err := connection.NewSqliteConnection("./")

	if err != nil {
		panic(err)
	}

	repository := repository.Newrepository(db)
	usecase := usecase.NewUsecase(repository)

}
