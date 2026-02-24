package main

import (
	"github.com/rafaelsouzaribeiro/exercicio-cotacao-mba-go-expert/internal/infra/web/client"
	"github.com/rafaelsouzaribeiro/exercicio-cotacao-mba-go-expert/internal/usecase"
)

func main() {

	usecase := usecase.UseCase{}
	client := client.NewClient(&usecase)
	err := client.CreateFile()

	if err != nil {
		panic(err)
	}
}
