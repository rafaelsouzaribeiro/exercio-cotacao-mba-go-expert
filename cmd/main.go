package main

import (
	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/infra/database/sqlite/connection"
	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/infra/database/sqlite/repository"
	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/infra/web/handlers"
	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/infra/web/server"
	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/usecase"
)

func main() {
	server := server.NewServer()
	db, err := connection.NewSqliteConnection("./cambio.db")

	if err != nil {
		panic(err)
	}

	repository := repository.NewRepository(db)
	usecase := usecase.NewUsecase(repository)
	handler := handlers.NewHandler(usecase)
	handler.Create()
	server.SetRoute(usecase)
	server.Start()

}
