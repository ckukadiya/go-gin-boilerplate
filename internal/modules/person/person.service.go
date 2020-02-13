package person

import (
	"context"
	"github.com/ckukadiya/go-gin-boilerplate/cmd/api/config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func New(db *mongo.Client, cfg *config.Configuration) *Service {
	return &Service{
		connection: db,
		database:   cfg.DB.Database,
		collection: "people",
	}
}

type Service struct {
	connection *mongo.Client
	database   string
	collection string
}

func (s *Service) List(c *gin.Context) *PersonListResponse {
	var people []*Person
	collection := s.connection.Database(s.database).Collection(s.collection)
	cur, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("Error on finding all documents ", err)
	}
	for cur.Next(context.TODO()) {
		var person Person
		err = cur.Decode(&person)
		if err != nil {
			log.Fatal("Error in decoding document ", err)
		}
		people = append(people, &person)
	}
	return &PersonListResponse{
		Error: false,
		Data:  people,
		Total: 0,
	}
}
