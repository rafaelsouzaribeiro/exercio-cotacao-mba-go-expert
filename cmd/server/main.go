package main

import (
	"github.com/rafaelsouzaribeiro/exercicio-cotacao-mba-go-expert/internal/infra/database/sqlite/connection"
	"github.com/rafaelsouzaribeiro/exercicio-cotacao-mba-go-expert/internal/infra/di"
	"github.com/rafaelsouzaribeiro/exercicio-cotacao-mba-go-expert/internal/infra/web/server"
)

func main() {
	server := server.NewServer()
	db, err := connection.NewSqliteConnection("./cambio.db")

	if err != nil {
		panic(err)
	}

	usecase := di.NewServerDI(db)
	server.SetRoute(usecase)
	server.Start()

}
