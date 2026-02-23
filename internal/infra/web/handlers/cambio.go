package handlers

import (
	"encoding/json"
	"net/http"
)

type Cambio struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func (h *Usecase) Cambio(w http.ResponseWriter, r *http.Request) {
	respo, err := http.Get("https://economia.awesomeapi.com.br/json/last/USD-BRL")

	if err != nil {
		http.Error(w, "Erro ao consumir a url", http.StatusInternalServerError)
	}
	defer respo.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respo.Body)
}
