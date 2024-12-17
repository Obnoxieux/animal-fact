package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	r := mux.NewRouter()

	/** ---------------- Routes ------------------ **/

	r.HandleFunc("/api/test", sendAPIStatus)

	err := http.ListenAndServe(":8090", r)
	if err != nil {
		log.Fatal(err)
	}
}

func sendAPIStatus(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := Response{
		Message: "Service available",
	}
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
