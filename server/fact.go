package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/animal-fact/service"
	"github.com/gorilla/mux"
	"github.com/pocketbase/dbx"
	"net/http"
	"strconv"
)

func RespondWithFact(db *dbx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s := service.DuckService{}

		vars := mux.Vars(r)
		id := vars["id"]
		if id == "" {
			errorBadRequest(w, r)
		}
		dbID, err := strconv.Atoi(id)
		if err != nil {
			errorBadRequest(w, r)
			return
		}
		fact, err := s.GetByID(db, dbID)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				errorNoResults(w, r)
				return
			} else {
				errorServerError(w, r)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(fact)
		if err != nil {
			errorServerError(w, r)
		}
	}
}

func RespondWithRandomFact(db *dbx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		s := service.DuckService{}
		fact, err := s.GetRandom(db)
		if err != nil {
			errorServerError(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(fact)
		if err != nil {
			errorServerError(w, r)
		}
	}
}
