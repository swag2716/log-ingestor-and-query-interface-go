package routes

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/log-ingestor-and-query-interface/handlers"
)

// Register HTTP endpoints
var Handlers = func(r *mux.Router, db *sql.DB) {
	r.HandleFunc("/ingest", func(w http.ResponseWriter, r *http.Request) {
		handlers.IngestHandler(w, r, db)
	}).Methods("POST")
}
