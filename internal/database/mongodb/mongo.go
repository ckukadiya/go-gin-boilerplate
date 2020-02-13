package mongodb

import (
	"context"
	"github.com/ckukadiya/go-gin-boilerplate/cmd/api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// New creates new database connection to a mongodb database
func New(cfg *config.Configuration) *mongo.Client {
	clientOptions := options.Client().ApplyURI(cfg.DB.Path)
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
