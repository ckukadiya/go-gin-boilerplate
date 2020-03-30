package base

import (
	"github.com/gin-gonic/gin"
	"go-gin-boilerplate/internal/database/mongodb"
)

type BaseControllerInterface interface {
	GetAll(*gin.Context, interface{})
	Fetch(*gin.Context, interface{})
	CreateOne(*gin.Context, interface{})
	Update(*gin.Context, interface{})
	Remove(*gin.Context, interface{})
	Prepare(interface{})
}

type BaseController struct {
	mongodb.MongoCrudService
}
