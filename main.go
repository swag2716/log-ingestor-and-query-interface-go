package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/log-ingestor-and-query-interface/routes"
)

var db *sql.DB

func main() {
	//Initialize PostgreSQL database

	initDB()

	// HTTP router
	r := mux.NewRouter()

	routes.Handlers(r, db)

	// Start the HTTP server on port 3000
	port := "3000"
	fmt.Println("Log Ingestor Server is running on port - ", port)
	log.Fatal(http.ListenAndServe(":3000", r))
}

// Connect to the PostgreSQL database
func initDB() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Println(err)
		return
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")

	connStr := fmt.Sprintf("host=%s port=%s user='%s' "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	// Create the logs table if it doesn't exist
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS logs (
			trace_id VARCHAR PRIMARY KEY,
			level VARCHAR,
			message VARCHAR,
			resource_id VARCHAR,
			timestamp TIMESTAMP,
			span_id VARCHAR,
			commit VARCHAR,
			parent_resource_id VARCHAR
		);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}
