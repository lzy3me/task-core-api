package routes

import (
	controller "task-core/controllers/board"

	"github.com/gofiber/fiber/v2"
)

func BoardRoute(app fiber.Router) {
	api := app.Group("/board")
	api.Post("/create", controller.Create)
	api.Get("/list", controller.ReadAll)
	api.Get("/:board_id", controller.ReadOne)
	api.Put("/:board_id", controller.Update)
	api.Delete("/:board_id", controller.Delete)
}
