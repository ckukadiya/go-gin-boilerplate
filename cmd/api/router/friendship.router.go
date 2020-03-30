package router

import (
	"github.com/gin-gonic/gin"
	"go-gin-boilerplate/internal/modules/friendship"
)

func NewFriendship(f *friendship.FriendshipController, r *gin.RouterGroup) {
	f.Prepare()

	personRoute := r.Group("/friendship")

	personRoute.GET("/list", f.GetAll)
	personRoute.GET("/get/:id", f.Get)
	personRoute.POST("/new", f.Post)
	personRoute.PUT("/update/:id", f.Put)
	personRoute.DELETE("/remove/:id", f.Delete)
}
