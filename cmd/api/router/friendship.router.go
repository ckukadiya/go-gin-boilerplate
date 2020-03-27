package router

import (
	"github.com/ckukadiya/go-gin-boilerplate/internal/modules/friendship"
	"github.com/gin-gonic/gin"
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
