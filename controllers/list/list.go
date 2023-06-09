package listController

import (
	"task-core/controllers/response"
	listEntity "task-core/models/entities/list"
	listRepository "task-core/models/repositories/list"

	"github.com/gofiber/fiber/v2"
)

func TemplateList(c *fiber.Ctx, res interface{}) (interface{}, error) {
	listBody := []interface{}{
		listEntity.Lists{
			IDBoard:  res,
			Name:     "To Do",
			Position: "1",
		},
		listEntity.Lists{
			IDBoard:  res,
			Name:     "Doing",
			Position: "2",
		},
		listEntity.Lists{
			IDBoard:  res,
			Name:     "Done",
			Position: "3",
		},
	}

	result, err := listRepository.InsertMany(c, listBody)
	return result, err
}

func Create(c *fiber.Ctx) error {
	body := new(listEntity.BodyCreate)
	if err := c.BodyParser(body); err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, "", nil)
	}

	rawBody := listEntity.Lists{
		IDBoard:  body.IDBoard,
		Name:     body.Name,
		Position: body.Position,
	}

	res, err := listRepository.Insert(c, rawBody)
	if err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return response.ResponseOK(c, fiber.StatusOK, res, "")
}

func Update(c *fiber.Ctx) error {
	return response.ResponseOK(c, fiber.StatusOK, "res", "")
}

func Archive(c *fiber.Ctx) error {
	return response.ResponseOK(c, fiber.StatusOK, "res", "")
}
