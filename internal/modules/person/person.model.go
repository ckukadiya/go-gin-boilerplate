package person

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	State  string `json:"state"`
}

type PersonModel interface {
	List(*gin.Context, bson.D, *options.FindOptions) (*PersonListResponse, error)
	Get(*gin.Context, bson.M) (*PersonGetResponse, error)
}

type PersonListResponse struct {
	Data  []*Person `json:"data"`
	Total int       `json:"total"`
}

type PersonGetResponse struct {
	Data *Person `json:"data"`
}
