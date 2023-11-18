package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/log-ingestor-and-query-interface/models"
)

var db *sql.DB

func IngestHandler(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v\n", r)
			debug.PrintStack()
		}
	}()
	var logEntry models.Log
	err := json.NewDecoder(r.Body).Decode(&logEntry)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoding log entry", http.StatusBadRequest)
		return
	}

	// Insert log entry into PostgreSQL
	_, err = db.Exec(`
		INSERT INTO logs 
		(trace_id, level, message, resource_id, timestamp, span_id, commit, parent_resource_id) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		ON CONFLICT (trace_id) DO NOTHING
	`,
		logEntry.TraceID,
		logEntry.Level,
		logEntry.Message,
		logEntry.ResourceID,
		logEntry.Timestamp,
		logEntry.SpanID,
		logEntry.Commit,
		logEntry.ParentResID,
	)
	if err != nil {
		http.Error(w, "Error inserting log entry into PostgreSQL", http.StatusInternalServerError)
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v\n", r)
			debug.PrintStack()
		}
	}()

	fmt.Fprintf(w, "Log ingested successfully!\n")
}
