package routes

import (
	bController "task-core/controllers/board"
	bcController "task-core/controllers/board_collection"

	"github.com/gofiber/fiber/v2"
)

func BoardRoute(app fiber.Router) {
	api := app.Group("/board")
	// Board
	api.Post("/create", bController.Create)
	api.Get("/list", bController.ReadAll)

	// Board Collection
	api.Get("/:boardId", bcController.ReadAll)
	api.Post("/:boardId/new", bcController.Create)
	api.Put("/edit/:boardCollectionId", bController.Delete) // Update board collection detail
	api.Put("/list/:boardCollectionId", bController.Delete) // Change list on board
}
