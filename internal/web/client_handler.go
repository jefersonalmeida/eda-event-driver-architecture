package web

import (
	"encoding/json"
	"github.com/jefersonalmeida/go-wallet/internal/usecase/create_client"
	"net/http"
)

type ClientHandler struct {
	CreateClientUseCase create_client.CreateClientUseCase
}

func NewClientHandler(createClientUseCase create_client.CreateClientUseCase) *ClientHandler {
	return &ClientHandler{
		CreateClientUseCase: createClientUseCase,
	}
}

func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var dto create_client.CreateClientInputDTO

	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	output, err := h.CreateClientUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusCreated)
}
