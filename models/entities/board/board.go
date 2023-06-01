package boardEnitity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Board struct {
	ID                primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	BelongToWorkspace primitive.ObjectID `json:"belongToWorkspace" bson:"belongToWorkspace"`
	BelongToUser      string             `json:"belongToUser" bson:"belongToUser"`
	Visibility        string             `json:"visibility" bson:"visibility"`
	List              []BoardList        `json:"list" bson:"list"`
	Permission        BoardPermission    `json:"permission" bson:"permission"`
	Labels            []BoardLabel       `json:"labels" bson:"labels"`
	UserMember        []BoardUserMember  `json:"userMember" bson:"userMember"`
}

type BoardLabel struct {
	Key   string `json:"key" bson:"key"`
	Name  string `json:"name" bson:"name"`
	Color string `json:"color" bson:"color"`
}

type BoardList struct {
	ID    string `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Order int    `json:"order" bson:"order"`
}

type BoardPermission struct {
	CommentLevel          int  `json:"commentLevel" bson:"commentLevel"`
	ManagePermissionLevel int  `json:"managePermissionLevel" bson:"managePermissionLevel"`
	WorkspaceMemberAccess bool `json:"workspaceMemberPermission" bson:"workspaceMemberPermission"`
}

type BoardUserMember struct {
	UserId primitive.ObjectID `json:"userId" bson:"userId"`
	Level  int                `json:"level" bson:"level"`
}
