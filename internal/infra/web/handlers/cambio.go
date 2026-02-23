package handlers

import (
	"encoding/json"
	"net/http"
)

func (h *Usecase) Cambio(w http.ResponseWriter, r *http.Request) {
	respo, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")

	if err != nil {
		http.Error(w, "Erro ao consumir a url", http.StatusInternalServerError)
	}
	defer respo.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respo.Body)
}
