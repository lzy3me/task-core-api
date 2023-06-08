package board

import (
	"log"
	listController "task-core/controllers/list"
	"task-core/controllers/response"
	boardEnitity "task-core/models/entities/board"
	boardRepository "task-core/models/repositories/board"
	listRepository "task-core/models/repositories/list"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Create(c *fiber.Ctx) error {
	body := new(boardEnitity.Body)
	if err := c.BodyParser(body); err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, "", nil)
	}

	orgId, _ := primitive.ObjectIDFromHex("6473b3bf609c1be218261551")
	rawBody := boardEnitity.Board{
		OrgID: orgId,
		Name:  body.Name,
	}

	result, err := boardRepository.Insert(c, rawBody)
	if err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	_, err = listController.TemplateList(c, result)
	if err != nil {
		log.Fatal(err)
	}

	return response.ResponseOK(c, fiber.StatusOK, result, "")
}

func ReadAll(c *fiber.Ctx) error {
	return response.ResponseOK(c, 200, "", "This is a read-all route")
}

func ReadOne(c *fiber.Ctx) error {
	var param boardEnitity.ParamCollection
	c.ParamsParser(&param)

	filter := bson.M{
		"_id": param.BoardID,
	}

	result, err := boardRepository.FindOne(c, filter)
	if err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return response.ResponseOK(c, fiber.StatusOK, result, "")
}

func Update(c *fiber.Ctx) error {
	return response.ResponseOK(c, 200, "", "This is a update route")
}

func Delete(c *fiber.Ctx) error {
	var param boardEnitity.ParamCollection
	c.ParamsParser(&param)

	filter := bson.M{
		"_id": param.BoardID,
	}

	filterList := bson.M{
		"idBoard": param.BoardID,
	}

	result, err := boardRepository.Delete(c, filter)
	if err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, err.Error(), nil)
	}
	_, err = listRepository.DeleteAll(c, filterList)
	if err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return response.ResponseOK(c, fiber.StatusOK, result, "")
}
