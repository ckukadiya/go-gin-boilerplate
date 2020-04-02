package person

import (
	"github.com/gin-gonic/gin"
	"go-gin-boilerplate/internal/modules/base"
	"net/http"
)

type PersonController struct {
	base.BaseController
}

func (this *PersonController) Prepare() {
	this.BindCollection("go-gin-boilerplate", "people")
}

func (this *PersonController) GetAll(c *gin.Context) {
	var result []Person
	if status := this.BaseController.GetAll(c, &result); status {
		response := PersonListResponse{
			Data:  result,
			Total: len(result),
		}
		c.JSON(http.StatusOK, response)
	}
}

func (this *PersonController) Get(c *gin.Context) {
	var result Person
	if status := this.Fetch(c, &result); status {
		response := PersonResponse{
			Data: result,
		}
		c.JSON(http.StatusOK, response)
	}
}

func (this *PersonController) Post(c *gin.Context) {
	var result Person
	if uid := this.CreateOne(c, &result); uid != nil {
		c.JSON(http.StatusOK, uid)
	}
}

func (this *PersonController) Put(c *gin.Context) {
	var result Person
	if status := this.Update(c, &result); status {
		response := PersonResponse{
			Data: result,
		}
		c.JSON(http.StatusOK, response)
	}
}

func (this *PersonController) Delete(c *gin.Context) {
	var result Person
	if status := this.Remove(c, &result); status {
		c.JSON(http.StatusOK, "Delete successfully!!!")
	}
}
