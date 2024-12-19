package server

import (
	"encoding/json"
	"github.com/animal-fact/model"
	"net/http"
)

func SendAPIStatus(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := model.SimpleResponse{
		Message: "Service available",
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
