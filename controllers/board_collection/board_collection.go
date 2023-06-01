package boardcollectionController

import (
	"log"
	"task-core/controllers/response"
	boardcollectionEntity "task-core/models/entities/board_collection"
	boardRepository "task-core/models/repositories/board"
	boardcollectionRepository "task-core/models/repositories/board_collection"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(c *fiber.Ctx) error {
	body := new(boardcollectionEntity.BodyCreate)
	var param boardcollectionEntity.ParamCollection

	if err := c.BodyParser(body); err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, "", nil)
	}
	c.ParamsParser(&param)

	rawBody := boardcollectionEntity.BoardCollection{
		BelongToBoard: param.BoardID,
		BelongToList:  body.BelongToList,
		Name:          body.Name,
	}

	result, err := boardcollectionRepository.Insert(c, rawBody)
	if err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return response.ResponseOK(c, fiber.StatusOK, result, "")
}

func ReadAll(c *fiber.Ctx) error {
	var query boardcollectionEntity.QueryCollection
	var param boardcollectionEntity.ParamCollection
	c.QueryParser(&query)
	c.ParamsParser(&param)

	filterB := bson.M{
		"_id": param.BoardID,
	}

	filterC := bson.M{
		"belongToBoard": param.BoardID,
		// "deleteAt": nil,
		// "isArchive": false,
	}

	if query.User != "" {
		userObjId, _ := primitive.ObjectIDFromHex(query.User)
		filterB["watchUser"] = userObjId
		filterB["assignUser"] = userObjId
	}

	if query.Name != "" {
		filterB["name"] = query.Name
	}

	match := bson.D{primitive.E{Key: "$match", Value: filterC}}
	pipeline := mongo.Pipeline{match}
	board, errB := boardRepository.FindOne(c, filterB)
	rows, errC := boardcollectionRepository.Find(c, pipeline)

	if errB != nil {
		log.Println("err query board /board/:boardId", errB)
	}
	if errC != nil {
		log.Println("err query collection /board/:boardId", errC)
	}

	data := boardcollectionEntity.ResponseList{
		Board: board,
		Rows:  rows,
	}

	return response.ResponseOK(c, fiber.StatusOK, data, "")
}

func ReadOne(c *fiber.Ctx) error {
	return response.ResponseOK(c, 200, "", "This is a read-one route")
}

func Update(c *fiber.Ctx) error {
	return response.ResponseOK(c, 200, "", "This is a update route")
}

func Delete(c *fiber.Ctx) error {
	return response.ResponseOK(c, 200, "", "This is a delete route")
}
