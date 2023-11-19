package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/log-ingestor-and-query-interface/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandlerLog(w http.ResponseWriter, r *http.Request) {
	var logEntry models.Log
	err := json.NewDecoder(r.Body).Decode(&logEntry)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error decoding log entry", http.StatusBadRequest)
		return
	}

	PostgreSQLIngest(logEntry)

	shardDB := selectMongoDBShard()

	MongoIngest(shardDB, logEntry)

	fmt.Fprintf(w, "Log ingested successfully!\n")
}

func selectMongoDBShard() *mongo.Database {
	mongoMutex.Lock()
	defer mongoMutex.Unlock()

	// Simple round-robin strategy
	selectedDB := mongoDBs[0]
	mongoDBs = append(mongoDBs[1:], mongoDBs[0])

	return selectedDB
}
