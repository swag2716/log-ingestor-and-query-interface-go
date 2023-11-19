package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/log-ingestor-and-query-interface/handlers"
)

// Register HTTP endpoints
var Handlers = func(r *mux.Router) {
	r.HandleFunc("/ingest", func(w http.ResponseWriter, r *http.Request) {
		handlers.HandlerLog(w, r)
	}).Methods("POST")
}
