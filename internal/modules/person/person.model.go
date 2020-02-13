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
	Error   bool      `json:"error"`
	Data    []*Person `json:"data"`
	Total   int       `json:"total"`
	Message string    `json:"message"`
}
