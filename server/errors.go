package server

import (
	"encoding/json"
	"github.com/animal-fact/model"
	"net/http"
)

func errorBadRequest(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	response := model.SimpleResponse{
		Message: "Bad Request",
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func errorServerError(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	response := model.SimpleResponse{
		Message: "Bad Request",
	}
	_ = json.NewEncoder(w).Encode(response)
}
