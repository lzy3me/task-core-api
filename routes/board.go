package routes

import (
	bController "task-core/controllers/board"
	bcController "task-core/controllers/board_collection"
	lController "task-core/controllers/list"

	"github.com/gofiber/fiber/v2"
)

func BoardRoute(app fiber.Router) {
	api := app.Group("/board")
	// Board
	api.Post("/create", bController.Create)
	api.Get("/b", bController.ReadAll)
	api.Get("/b/:boardId", bController.ReadOne)
	api.Put("/b/:boardId", bController.Update)
	api.Delete("/b/:boardId", bController.Delete)

	// List
	api.Post("/l/create", lController.Create)
	api.Put("/l/:listId", lController.Update)
	api.Put("/l/a/:listId", lController.Archive)

	// Board Collection
	api.Get("/:boardId", bcController.ReadAll)
	api.Post("/:boardId/new", bcController.Create)
	api.Get("/detail/:boardCollectionId", bcController.Detail)
	api.Put("/edit/:boardCollectionId", bcController.Update)     // Update board collection detail
	api.Put("/list/:boardCollectionId", bcController.ChangeList) // Change list on board
}
