package boardcollectionEntity

import (
	Board "task-core/models/entities/board"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BoardCollection struct {
	ID            primitive.ObjectID   `json:"_id" bson:"_id,omitempty"`
	BelongToBoard primitive.ObjectID   `json:"belongToBoard" bson:"belongToBoard"`
	BelongToList  string               `json:"belongToList" bson:"belongToList"`
	Name          string               `json:"name" bson:"name"`
	Description   string               `json:"description" bson:"description"`
	WatchUsers    []primitive.ObjectID `json:"watchUsers" bson:"watchUsers"`
	AssignUsers   []primitive.ObjectID `json:"assignUsers" bson:"assignUsers"`
	DueDate       string               `json:"due_date" bson:"due_date"`
	Labels        []string             `json:"label" bson:"label"`
	Collection    []string             `json:"collection" bson:"collection"`
}

type BodyCreate struct {
	BelongToBoard primitive.ObjectID `json:"belongToBoard" bson:"belongToBoard"`
	BelongToList  string             `json:"belongToList" bson:"belongToList"`
	Name          string             `json:"name" bson:"name"`
}

type BodyEdit struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	WatchUsers  string `json:"watchUsers" bson:"watchUsers"`
	AssignUsers string `json:"assignUsers" bson:"assignUsers"`
	DueDate     string `json:"due_date" bson:"due_date"`
	Labels      string `json:"label" bson:"label"`
	Collection  string `json:"collection" bson:"collection"`
}

type QueryCollection struct {
	User string `json:"userId" bson:"userId"`
	Name string `json:"name" bson:"name"`
}

type BodyChangeList struct {
	ListID string `json:"listId", bson:"listId"`
}

type ParamCollection struct {
	BoardCollectionID primitive.ObjectID `json:"boardCollectionId" bson:"boardCollectionId"`
}

type ResponseList struct {
	Board Board.Board       `json:"board" bson:"board"`
	Rows  []BoardCollection `json:"collections" bson:"collections"`
}
