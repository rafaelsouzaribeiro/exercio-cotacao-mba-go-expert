package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rafaelsouzaribeiro/exercio-cotacao-mba-go-expert/internal/entity"
)

func (h *Usecase) Cambio(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")

	if err != nil {
		http.Error(w, "Erro ao consumir a url", http.StatusInternalServerError)
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
