package boardcollectionController

import (
	"log"
	"task-core/controllers/response"
	boardEnitity "task-core/models/entities/board"
	boardcollectionEntity "task-core/models/entities/board_collection"
	boardRepository "task-core/models/repositories/board"
	boardcollectionRepository "task-core/models/repositories/board_collection"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create new board collection
func Create(c *fiber.Ctx) error {
	body := new(boardcollectionEntity.BodyCreate)
	var param boardEnitity.ParamCollection

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

// Read all board collection on board target id
// return board data and board collections
func ReadAll(c *fiber.Ctx) error {
	var query boardcollectionEntity.QueryCollection
	var param boardEnitity.ParamCollection
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

func Detail(c *fiber.Ctx) error {
	var param boardcollectionEntity.ParamCollection

	c.ParamsParser(&param)

	filter := bson.M{
		"_id": param.BoardCollectionID,
	}
	boardCollection, err := boardcollectionRepository.FindOne(c, filter)
	if err != nil {
		log.Println("err query collection /board/detail/:boardCollectionId", err)
	}

	return response.ResponseOK(c, 200, boardCollection, "")
}

func Update(c *fiber.Ctx) error {
	body := new(boardcollectionEntity.BodyEdit)
	var param boardcollectionEntity.ParamCollection

	if err := c.BodyParser(body); err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, "", nil)
	}
	c.ParamsParser(&param)

	filter := bson.M{
		"_id": param.BoardCollectionID,
		// "deleteAt": nil,
	}

	// now := time.Now()
	rawBody := bson.M{}

	if body.Name != "" {
		rawBody["name"] = body.Name
	}

	if body.Description != "" {
		rawBody["description"] = body.Description
	}

	if body.DueDate != "" {
		rawBody["due_date"] = body.DueDate
	}

	if body.Labels != "" {
		rawBody["label"] = body.Labels
	}

	if rawBody == nil {
		return response.ResponseOK(c, fiber.StatusOK, "", "")
	}

	update := bson.M{
		"$set": rawBody,
	}

	result, err := boardcollectionRepository.Update(c, filter, update, false)
	if err != nil {
		log.Println("err /board/edit/:boardCollectionId", err)
		return response.ResponseError(c, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message, "")
	}

	return response.ResponseOK(c, fiber.StatusOK, result, "")
}

func ChangeList(c *fiber.Ctx) error {
	body := new(boardcollectionEntity.BodyChangeList)
	var param boardcollectionEntity.ParamCollection

	if err := c.BodyParser(body); err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, "", nil)
	}
	c.ParamsParser(&param)

	filter := bson.M{
		"_id": param.BoardCollectionID,
		// "deleteAt": nil,
	}

	// now := time.Now()
	rawBody := bson.M{
		"$set": bson.M{
			"belongToList": body.ListID,
			// "updateAt": now,
		},
	}

	result, err := boardcollectionRepository.Update(c, filter, rawBody, false)
	if err != nil {
		log.Println("err /board/list/:boardCollectionId", err)
		return response.ResponseError(c, fiber.StatusInternalServerError, fiber.ErrInternalServerError.Message, "")
	}

	return response.ResponseOK(c, fiber.StatusOK, result, "")
}

func Delete(c *fiber.Ctx) error {
	return response.ResponseOK(c, 200, "", "This is a delete route")
}
