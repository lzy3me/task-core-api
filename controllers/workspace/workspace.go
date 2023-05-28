package workspace

import (
	"task-core/controllers/response"
	workspaceEntitiy "task-core/models/entities/workspace"
	workspaceRepository "task-core/models/repositories/workspace"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	body := new(workspaceEntitiy.BodyCreateWorkspace)

	if err := c.BodyParser(body); err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, "", nil)
	}

	rawBody := workspaceEntitiy.Workspace{
		Type:        body.Type,
		Name:        body.Name,
		Description: body.Description,
		Visibility:  body.Visibility,
	}

	result, err := workspaceRepository.Insert(c, rawBody)
	if err != nil {
		return response.ResponseError(c, fiber.StatusBadRequest, err.Error(), nil)
	}

	return response.ResponseOK(c, fiber.StatusOK, result, "")
}

func ReadAll(c *fiber.Ctx) error {
	return response.ResponseOK(c, 200, "", "This is a read-all route")
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
