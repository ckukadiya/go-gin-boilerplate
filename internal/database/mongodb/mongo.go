package mongodb

import (
	"context"
	"github.com/ckukadiya/go-gin-boilerplate/cmd/api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
)

type MongoService struct {
	Collection *mongo.Collection
}

var db *mongo.Client

// Singleton
var dial sync.Once

func (c *MongoService) BindCollection(database string, collection string) {
	c.Collection = db.Database(database).Collection(collection)
}

// New creates new database connection to a mongodb database
// New creates new database connection to a mongodb database
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

func NewMongoDB() {
	path := config.GetConfig().DB.Path
	dial.Do(func() {
		db = new(path)
		log.Println("MongoDB Connected at " + path)
	})
}

func GetMongoDBClient() *mongo.Client {
	return db
}
