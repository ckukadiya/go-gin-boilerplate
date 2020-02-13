package router

import (
	"github.com/ckukadiya/go-gin-boilerplate/internal/modules/person"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Person struct {
	service *person.Service
}

func NewPerson(s *person.Service, r *gin.RouterGroup) {
	p := Person{service: s}
	personRoute := r.Group("/person")

	personRoute.GET("/list", p.list)
}

func (p *Person) list(c *gin.Context) {
	res := p.service.List(c)
	c.JSON(http.StatusOK, res)
}
