package router

import (
	"github.com/ckukadiya/go-gin-boilerplate/cmd/api/middelware"
	apperror "github.com/ckukadiya/go-gin-boilerplate/internal/error"
	"github.com/ckukadiya/go-gin-boilerplate/internal/modules/person"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	res, err := p.service.List(c, bson.M{})
	if err != nil {
		apperror.Response(c, err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (p *Person) get(c *gin.Context) {
	id, err := middelware.ID(c)
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
