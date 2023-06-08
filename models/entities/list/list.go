package listEntity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Lists struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	IDBoard  any                `json:"idBoard" bson:"idBoard"`
	Name     string             `json:"name" bson:"name"`
	Position string             `json:"pos" bson:"pos"`
}

type BodyCreate struct {
	IDBoard  primitive.ObjectID `json:"idBoard" bson:"idBoard"`
	Name     string             `json:"name" bson:"name"`
	Position string             `json:"pos" bson:"pos"`
}
