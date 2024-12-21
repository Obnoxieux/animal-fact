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

	/** ---------------- Static Assets ------------------ **/
	fs := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	/** ---------------- Routes ------------------ **/

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		server.RenderTemplate(w, "index", nil)
	})

	r.HandleFunc("/impressum", func(w http.ResponseWriter, r *http.Request) {
		server.RenderTemplate(w, "impressum", nil)
	})

	r.HandleFunc("/privacy", func(w http.ResponseWriter, r *http.Request) {
		server.RenderTemplate(w, "privacy", nil)
	})

	r.HandleFunc("/api/test", server.SendAPIStatus)

	r.HandleFunc("/facts/duck/random", server.RespondWithRandomFact(db)).Methods(http.MethodGet)
	r.HandleFunc("/facts/duck/{id:[0-9]+}", server.RespondWithFact(db)).Methods(http.MethodGet)

	err = http.ListenAndServe(":8090", r)
	if err != nil {
		log.Fatal(err)
	}
}
