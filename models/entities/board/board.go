package boardEnitity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Board struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	CreatorID     primitive.ObjectID `json:"creatorId" bson:"creatorId"`
	OrgID         primitive.ObjectID `json:"organizationId" bson:"organizationId"`
	Name          string             `json:"name" bson:"name"`
	IsDeactivate  bool               `json:"isDeactivate" bson:"isDeactivate"`
	DeactivateDue string             `json:"dueDeactivate" bson:"dueDeactivate"`
}

type ParamCollection struct {
	BoardID primitive.ObjectID `json:"boardId" bson:"boardId"`
}

type Body struct {
	CreatorID primitive.ObjectID `json:"creatorId" bson:"creatorId"`
	OrgID     primitive.ObjectID `json:"organizationId" bson:"organizationId"`
	Name      string             `json:"name" bson:"name"`
}

type SoftDeleteBody struct {
	IsDeactivate  bool   `json:"isDeactivate" bson:"isDeactivate"`
	DeactivateDue string `json:"dueDeactivate" bson:"dueDeactivate"`
}
