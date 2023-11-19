package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/log-ingestor-and-query-interface/handlers"
	"github.com/log-ingestor-and-query-interface/routes"
)

var db *sql.DB

func main() {
	//Initialize PostgreSQL database

	handlers.InitDB()
	handlers.InitMongoDB()

	// HTTP router
	r := mux.NewRouter()

	routes.Handlers(r)

	// Start the HTTP server on port 3000
	port := "3000"
	fmt.Println("Log Ingestor Server is running on port - ", port)
	log.Fatal(http.ListenAndServe(":3000", r))

	// simulateLogs()
}

// func simulateLogs() {
// 	for {
// 		logEntry := LogEntry{
// 			Timestamp: time.Now(),
// 			Message:   "Sample log message",
// 		}
// 		handleLogEntry(logEntry)
// 		time.Sleep(time.Second)
// 	}
// }
