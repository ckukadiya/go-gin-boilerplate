package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FilterOptions struct {
	Filter      bson.D
	FindOptions *options.FindOptions
}
