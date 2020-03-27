package mongodb

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

type MongoCrudService struct {
	MongoService
}

func (this *MongoCrudService) GetAllResults(filterOptions FilterOptions, res interface{}) (err error) {
	if this.Collection == nil {
		log.Println("Collection is not defined, Please override Prepare and Bind collection ")
		return errors.New("Collection is not defined")
	}
	cur, err := this.Collection.Find(context.TODO(), filterOptions.Filter, filterOptions.FindOptions)
	if err != nil {
		return
	}
	err = cur.All(context.TODO(), res)
	if err != nil {
		return
	}
	return
}

func (this *MongoCrudService) GetOne(filter bson.D, res interface{}) error {
	if this.Collection == nil {
		log.Println("Collection is not defined, Please override Prepare and Bind collection ")
		return errors.New("Collection is not defined")
	}
	err := this.Collection.FindOne(context.TODO(), filter).Decode(res)
	if err != nil {
		return err
	}
	return nil
}

func (this *MongoCrudService) InsertOne(data interface{}) (interface{}, error) {
	if this.Collection == nil {
		log.Println("Collection is not defined, Please override Prepare and Bind collection ")
		return nil, errors.New("Collection is not defined")
	}
	result, err := this.Collection.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (this *MongoCrudService) FindOneAndUpdate(filter bson.D, data interface{}) (interface{}, error) {
	if this.Collection == nil {
		log.Println("Collection is not defined, Please override Prepare and Bind collection ")
		return nil, errors.New("Collection is not defined")
	}
	return this.Collection.FindOneAndUpdate(context.TODO(), filter, bson.D{{Key: "$set", Value: data}}).Decode(&data), nil
}

func (this *MongoCrudService) FindOneAndRemove(filter bson.D, data interface{}) error {
	if this.Collection == nil {
		log.Println("Collection is not defined, Please override Prepare and Bind collection ")
		return errors.New("Collection is not defined")
	}
	this.Collection.FindOneAndDelete(context.TODO(), filter).Decode(&data)
	return nil
}
