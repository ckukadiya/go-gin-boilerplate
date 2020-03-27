package base

import (
	request_parser "github.com/ckukadiya/go-gin-boilerplate/cmd/api/request-parser"
	"github.com/ckukadiya/go-gin-boilerplate/internal/database/mongodb"
	apperror "github.com/ckukadiya/go-gin-boilerplate/internal/error"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (this *BaseController) GetAll(c *gin.Context, results interface{}) bool {
	pagination, _ := request_parser.Paginate(c)
	findOptions := options.Find()
	findOptions.SetLimit(pagination.Limit)
	findOptions.Skip = &pagination.Skip
	filterOptions := mongodb.FilterOptions{bson.D{}, findOptions}
	err := this.GetAllResults(filterOptions, results)
	if err != nil {
		apperror.Response(c, err)
		return false
	}
	return true
}

func (this *BaseController) Fetch(c *gin.Context, result interface{}) bool {
	uid, err := request_parser.ID(c)
	if err != nil {
		apperror.Response(c, err)
		return false
	}
	if uid != "" {
		objectId, _ := primitive.ObjectIDFromHex(uid)
		err := this.GetOne(bson.D{{"_id", objectId}}, result)
		if err != nil {
			apperror.Response(c, err)
			return false
		}
	}
	return true
}

func (this *BaseController) CreateOne(c *gin.Context, data interface{}) interface{} {
	err := c.Bind(data)
	if err == nil {
		uid, err := this.InsertOne(data)
		if err != nil {
			apperror.Response(c, err)
			return nil
		}
		return uid
	}
	apperror.Response(c, err)
	return nil
}

func (this *BaseController) Update(c *gin.Context, data interface{}) bool {
	uid, err := request_parser.ID(c)
	if err != nil {
		apperror.Response(c, err)
		return false
	}
	if uid != "" {
		if err := c.ShouldBind(data); err != nil {
			objectId, _ := primitive.ObjectIDFromHex(uid)
			if _, errorOnUpdate := this.FindOneAndUpdate(bson.D{{"_id", objectId}}, data); errorOnUpdate != nil {
				apperror.Response(c, errorOnUpdate)
				return false
			}
			return true
		}
		apperror.Response(c, err)
	}
	return false
}

func (this *BaseController) Remove(c *gin.Context, data interface{}) bool {
	uid, err := request_parser.ID(c)
	if err != nil {
		apperror.Response(c, err)
		return false
	}
	objectId, _ := primitive.ObjectIDFromHex(uid)
	if err = this.FindOneAndRemove(bson.D{{"_id", objectId}}, data); err != nil {
		apperror.Response(c, err)
		return false
	}
	return true
}
