package cardEntity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Card struct {
	ID          primitive.ObjectID   `json:"_id" bson:"_id,omitempty"`
	BoardID     primitive.ObjectID   `json:"idBoard" bson:"idBoard"`
	ListID      primitive.ObjectID   `json:"idList" bson:"idList"`
	Members     []primitive.ObjectID `json:"members" bson:"members"`
	Position    string               `json:"pos" bson:"pos"`
	Name        string               `json:"name" bson:"name"`
	Description string               `json:"desc" bson:"desc"`
	Cover       CardCover            `json:"cover" bson:"cover"`
	Due         primitive.DateTime   `json:"due" bson:"due"`
	DueComplete bool                 `json:"dueComplete" bson:"dueComplete"`
	DueReminder primitive.DateTime   `json:"dueReminder" bson:"dueReminder"`
	IsArchive   bool                 `json:"isArchive" bson:"isArchive"`
}

type CardCover struct {
	Attachment string `json:"attactment" bson:"attactment"`
	Color      string `json:"color" bson:"color"`
	Brightness string `json:"brightness" bson:"brightness"`
	Size       string `json:"size" bson:"size"`
}

type BodyCreate struct {
	BoardID  primitive.ObjectID `json:"idBoard" bson:"idBoard"`
	ListID   primitive.ObjectID `json:"idList" bson:"idList"`
	Position string             `json:"pos" bson:"pos"`
	Name     string             `json:"name" bson:"name"`
}

type BodyEdit struct {
	ListID      primitive.ObjectID   `json:"idList" bson:"idList"`
	Position    string               `json:"pos" bson:"pos"`
	Name        string               `json:"name" bson:"name"`
	Members     []primitive.ObjectID `json:"members" bson:"members"`
	Description string               `json:"desc" bson:"desc"`
	Cover       CardCover            `json:"cover" bson:"cover"`
	Due         primitive.DateTime   `json:"due" bson:"due"`
	DueComplete bool                 `json:"dueComplete" bson:"dueComplete"`
	DueReminder primitive.DateTime   `json:"dueReminder" bson:"dueReminder"`
	IsArchive   bool                 `json:"isArchive" bson:"isArchive"`
}
