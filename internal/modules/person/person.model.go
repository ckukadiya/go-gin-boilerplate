package person

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Person represents the details of single person
//
// swagger:model
type Person struct {
	Id     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name"`
	Age    int                `json:"age"`
	Gender string             `json:"gender"`
	State  string             `json:"state"`
}

// PersonListResponse represents the response of all people
//
// swagger:model
type PersonListResponse struct {
	Data  []Person `json:"data"`
	Total int      `json:"total"`
}

// PersonResponse represents the response of person details
//
// swagger:model
type PersonResponse struct {
	Data Person `json:"data"`
}
