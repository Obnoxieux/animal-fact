package main

import (
	"github.com/animal-fact/server"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/pocketbase/dbx"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"
)

// / Loads environment.
//   - Locally: from .env file
//   - Production: from container env
func init() {
	_ = gotenv.Load()
	if os.Getenv("PG_DSN") == "" {
		log.Fatal("PG_DSN not set")
	}
}

func main() {
	dsn := os.Getenv("PG_DSN")

	db, err := dbx.MustOpen("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	/** ---------------- Routes ------------------ **/

	r.HandleFunc("/api/test", server.SendAPIStatus)

	r.HandleFunc("/facts/duck/{id}", server.RespondWithFact(db)).Methods(http.MethodGet)

	err = http.ListenAndServe(":8090", r)
	if err != nil {
		log.Fatal(err)
	}
}
