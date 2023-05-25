package routes

import (
	controller "task-core/controllers/workspace"

	"github.com/gofiber/fiber/v2"
)

func WorkspaceRoute(app fiber.Router) {
	api := app.Group("/workspace")
	api.Post("/create", controller.Create)
	api.Get("/list", controller.ReadAll)
	api.Get("/:workspace_id", controller.ReadOne)
	api.Put("/:workspace_id", controller.Update)
	api.Delete("/:workspace_id", controller.Delete)
}
