package workspaceEntitiy

import (
	"task-core/models/entities"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BodyCreateWorkspace struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name        string             `json:"workspaceName" bson:"workspaceName"`
	Type        string             `json:"workspaceType" bson:"workspaceType"`
	Description string             `json:"description" bson:"description"`
	Visibility  string             `json:"workspaceVisibility" bson:"workspaceVisibility"`
}

type QueryList struct {
	Name string `json:"workspaceName" bson:"workspaceName"`
	Type string `json:"workspaceType" bson:"workspaceType"`
	entities.PaginationRequests
}

type ResponseList struct {
	Rows []Workspace `json:"rows"`
}

type Workspace struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name           string             `json:"workspaceName" bson:"workspaceName"`
	Type           string             `json:"workspaceType" bson:"workspaceType"`
	Description    string             `json:"description" bson:"description"`
	Visibility     string             `json:"workspaceVisibility" bson:"workspaceVisibility"`
	Member         []UserMember       `json:"workspaceMember" bson:"workspaceMember"`
	PermissionList []PermissionList   `json:"workspacePermission" bson:"workspacePermission"`
}

type UserMember struct {
	UserID primitive.ObjectID `json:"userId" bson:"userId"`
	RoleID string             `json:"roleId" bson:"roleId"`
}

type PermissionList struct {
	ID         string     `json:"id" bson:"id"`
	RoleName   string     `json:"roleName" bson:"roleName"`
	Permission Permission `json:"permission" bson:"permission"`
}

type Permission struct {
	CanCreate Group `json:"canCreate" bson:"canCreate"`
	CanDrop   Group `json:"canDrop" bson:"canDrop"`
	CanInvite bool  `json:"canInvite" bson:"canInvite"`
	CanShare  bool  `json:"canShare" bson:"canShare"`
}

type Group struct {
	Public    bool `json:"public" bson:"public"`
	Private   bool `json:"private" bson:"private"`
	Workspace bool `json:"workspace" bson:"workspace"`
}
