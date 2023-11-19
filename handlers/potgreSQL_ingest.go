package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/log-ingestor-and-query-interface/models"
)

var db *sql.DB

// Connect to the PostgreSQL database
func InitDB() {
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

func PostgreSQLIngest(logEntry models.Log) {

	// Insert log entry into PostgreSQL
	_, err := db.Exec(`
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
		log.Println("Error storing log in MongoDB:", err)
		return
	}

}
