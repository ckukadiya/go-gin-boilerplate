package router

import (
	request_parser "github.com/ckukadiya/go-gin-boilerplate/cmd/api/request-parser"
	apperror "github.com/ckukadiya/go-gin-boilerplate/internal/error"
	"github.com/ckukadiya/go-gin-boilerplate/internal/modules/person"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
)

type Person struct {
	service *person.Service
}

func NewPerson(s *person.Service, r *gin.RouterGroup) {
	p := Person{service: s}
	personRoute := r.Group("/person")

	personRoute.GET("/list", p.list)
	personRoute.GET("/get/:id", p.get)
}

func (p *Person) list(c *gin.Context) {
	pagination, _ := request_parser.Paginate(c)
	findOptions := options.Find()
	findOptions.SetLimit(pagination.Limit)
	findOptions.Skip = &pagination.Skip
	res, err := p.service.List(c, bson.D{{}}, findOptions)
	if err != nil {
		apperror.Response(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (p *Person) get(c *gin.Context) {
	id, err := request_parser.ID(c)
	if err != nil {
		apperror.Response(c, err)
		return
	}
	objectID, _ := primitive.ObjectIDFromHex(id)
	res, err := p.service.Get(c, bson.M{"_id": objectID})
	if err != nil {
		apperror.Response(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
