package person

import (
	"context"
	"github.com/ckukadiya/go-gin-boilerplate/cmd/api/config"
	apperror "github.com/ckukadiya/go-gin-boilerplate/internal/error"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
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

func (s *Service) List(c *gin.Context, filter bson.M) (*PersonListResponse, error) {
	var people []*Person
	collection := s.connection.Database(s.database).Collection(s.collection)
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		log.Println("Error on finding all documents ", err.Error())
		return nil, apperror.BadRequest
	}
	for cur.Next(context.TODO()) {
		var person Person
		err = cur.Decode(&person)
		if err != nil {
			log.Println("Error in decoding document ", err.Error())
		}
		people = append(people, &person)
	}
	return &PersonListResponse{
		Data:  people,
		Total: 0,
	}, nil
}

func (s *Service) Get(c *gin.Context, filter bson.M) (*PersonGetResponse, error) {
	var person Person
	collection := s.connection.Database(s.database).Collection(s.collection)
	err := collection.FindOne(context.TODO(), filter).Decode(&person)
	if err != nil {
		return nil, apperror.New(http.StatusOK, err.Error())
	}
	return &PersonGetResponse{
		Data: &person,
	}, nil
}
