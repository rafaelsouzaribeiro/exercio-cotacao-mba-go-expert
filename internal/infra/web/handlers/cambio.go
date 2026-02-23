package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/entity"
)

func (h *Usecase) Cambio(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)

	if err != nil {
		http.Error(w, "Erro ao criar a requisição", http.StatusInternalServerError)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		http.Error(w, "Erro ao consumir a url", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var cambio entity.Cambio
	if err := json.NewDecoder(resp.Body).Decode(&cambio); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cambio)
}
