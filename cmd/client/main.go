package main

import (
	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/infra/web/client"
	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/usecase"
)

func main() {

	usecase := usecase.UseCase{}
	client := client.NewClient(&usecase)
	err := client.CreateFile()

	if err != nil {
		panic(err)
	}
}
