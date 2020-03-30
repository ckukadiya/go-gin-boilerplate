package friendship

import (
	"github.com/gin-gonic/gin"
	"go-gin-boilerplate/internal/modules/base"
	"net/http"
)

type FriendshipController struct {
	base.BaseController
}

func (this *FriendshipController) Prepare() {
	this.BindCollection("mongo-tigergraph", "friendships")
}

func (this *FriendshipController) GetAll(c *gin.Context) {
	var result []Friendship
	if status := this.BaseController.GetAll(c, &result); status {
		c.JSON(http.StatusOK, result)
	}
}

func (this *FriendshipController) Get(c *gin.Context) {
	var result Friendship
	if status := this.Fetch(c, &result); status {
		c.JSON(http.StatusOK, result)
	}
}

func (this *FriendshipController) Post(c *gin.Context) {
	var result Friendship
	if uid := this.CreateOne(c, &result); uid != nil {
		c.JSON(http.StatusOK, uid)
	}
}

func (this *FriendshipController) Put(c *gin.Context) {
	var result Friendship
	if status := this.Update(c, &result); status {
		c.JSON(http.StatusOK, result)
	}
}

func (this *FriendshipController) Delete(c *gin.Context) {
	var result Friendship
	if status := this.Remove(c, &result); status {
		c.JSON(http.StatusOK, "Delete successfully!!!")
	}
}
