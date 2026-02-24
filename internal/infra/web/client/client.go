package client

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/entity"
)

func (c *Client) CreateFile() error {

	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/cotacao", nil)

	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	var cambio entity.Cambio
	if err := json.NewDecoder(resp.Body).Decode(&cambio); err != nil {
		return err
	}

	err = c.usecase.CreateFile(cambio.USDBRL.Bid)

	if err != nil {
		return err
	}

	return nil
}
