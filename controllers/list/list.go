package listController

import (
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
