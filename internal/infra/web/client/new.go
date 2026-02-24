package client

import "github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/usecase"

type Client struct {
	usecase *usecase.UseCase
}

func NewClient(usecase *usecase.UseCase) *Client {
	return &Client{
		usecase: usecase,
	}
}
