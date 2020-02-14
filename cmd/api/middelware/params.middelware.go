package middelware

import (
	apperror "github.com/ckukadiya/go-gin-boilerplate/internal/error"
	"github.com/gin-gonic/gin"
)

const defaultLimit = 100
const maxLimit = 1000

// Pagination contains pagination request
type Pagination struct {
	Limit  int `form:"limit"`
	Page   int `form:"page" binding:"min=0"`
	Offset int `json:"-"`
}

// Paginate validates pagination requests
func Paginate(c *gin.Context) (*Pagination, error) {
	p := new(Pagination)
	if err := c.ShouldBindQuery(p); err != nil {
		apperror.Response(c, err)
		return nil, err
	}
	if p.Limit < 1 {
		p.Limit = defaultLimit
	}
	if p.Limit > 1000 {
		p.Limit = maxLimit
	}
	p.Offset = p.Limit * p.Page
	return p, nil
}

// ID returns id url parameter.
// In case of conversion error to int, request will be aborted with StatusBadRequest.
func ID(c *gin.Context) (string, error) {
	id := c.Param("id")
	if id == "" {
		e := apperror.BadRequest
		e.Message = "Param Id not found in request"
		return "", e
	}
	return id, nil
}
