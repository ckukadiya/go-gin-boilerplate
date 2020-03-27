package friendship

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Friendship struct {
	Id      primitive.ObjectID `bson:"-" json:"-"`
	Person1 string             `json:"person1" form:"person1" binding:"required"`
	Person2 string             `json:"person2" form:"person2" binding:"required"`
	Date    time.Time          `bson:"connect_day" form:"connect_day" json:"connect_day" binding:"required"`
}
