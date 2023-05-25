package board

import (
	"task-core/controllers/response"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	return response.ResponseOK(c, 200, "", "This is a create route")
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
