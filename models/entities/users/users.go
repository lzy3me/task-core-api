package usersEntity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Username string             `json:"username" bson:"username"`
	Email    string             `json:"email" bson:"email"`
	Hash     string             `json:"hash" bson:"hash"`
	Salt     string             `json:"salt" bson:"salt"`
}

type UserProfile struct {
	ID         primitive.ObjectID
	UserID     primitive.ObjectID
	FirstName  string
	MiddleName string
	LastName   string
	Avatar     string
}
