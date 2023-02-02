package mongodb

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct {
	Collection *mongo.Collection
}

var db *mongo.Client

// Singleton
var dial sync.Once

func (c *MongoService) BindCollection(database, collection string) {
	c.Collection = db.Database(database).Collection(collection)
}

// New create new database connection to a mongodb database
func new(path string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(path)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return client
}

// Create new mongodb
func NewMongoDB(path string) {
	dial.Do(func() {
		db = new(path)
		log.Println("MongoDB Connected at " + path)
	})
}

// Get MongoDB Client
func GetMongoDBClient() *mongo.Client {
	return db
}
