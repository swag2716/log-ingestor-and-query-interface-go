package handlers

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/log-ingestor-and-query-interface/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoMutex sync.Mutex
	mongoDBs   []*mongo.Database
	shardURL   = "mongodb://localhost:27017/log-ingestion"
)

func InitMongoDB() {

	mongoDBs = connectMongoDBShards()
	defer func() {
		mongoMutex.Lock()
		defer mongoMutex.Unlock()
		for _, db := range mongoDBs {
			if err := db.Client().Disconnect(context.Background()); err != nil {
				log.Fatal(err)
			}
		}
	}()
}

func connectMongoDBShards() []*mongo.Database {
	// Simulate connecting to multiple MongoDB shards (replica sets)
	numShards := 3
	var databases []*mongo.Database

	for i := 0; i < numShards; i++ {
		shardURL := fmt.Sprintf("%s%d", shardURL, i)
		clientOptions := options.Client().ApplyURI(shardURL)
		client, err := mongo.Connect(context.Background(), clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		if err := client.Ping(context.Background(), nil); err != nil {
			log.Fatal(err)
		}

		databases = append(databases, client.Database("logs"))
	}

	return databases
}

func MongoIngest(db *mongo.Database, logEntry models.Log) {

	// Choose a MongoDB shard based on a simple round-robin strategy

	out, err := db.Collection("logs").InsertOne(context.Background(), logEntry)
	fmt.Println(out)
	if err != nil {
		log.Println("Error storing log in MongoDB:", err)
	}

	// w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "Log ingested successfully!\n")

}
