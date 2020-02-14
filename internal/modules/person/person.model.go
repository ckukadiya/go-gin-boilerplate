package person

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
	State  string `json:"state"`
}

type PersonModel interface {
	List(*gin.Context) *PersonListResponse
}

type PersonListResponse struct {
	Data  []*Person `json:"data"`
	Total int       `json:"total"`
}

type PersonGetResponse struct {
	Data *Person `json:"data"`
}
