package routes

import "github.com/gofiber/fiber/v2"

func RootRoute(app fiber.Router) {
	WorkspaceRoute(app)
	BoardRoute(app)
	TaskRoute(app)
}
