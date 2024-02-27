package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Netflix struct {
	Id      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Movie   string `json:"_id,omitempty"`
	Watched bool `json:"_id,omitempty"`
}
