package main

import (
	"errors"
	"fmt"
	"task-core/routes"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recoverFiber "github.com/gofiber/fiber/v2/middleware/recover"
)

func Handle(c *fiber.Ctx) error {
	defer func() {
		if err := recover(); err != nil {

			errs := fmt.Sprintf("%s", err)
			errr := errors.New(errs)
			sentry.CaptureException(errr)
		}
	}()
	return c.Next()
}

func ErrorHandler() fiber.Handler {
	return Handle
}

func main() {
	port := ":4500"

	app := fiber.New(fiber.Config{
		ReadBufferSize: 100 * 1024 * 1024,
		BodyLimit:      100 * 1024 * 1024,
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			if code == fiber.StatusInternalServerError {
				defer sentry.CaptureException(err)
				// res := entity.ResponseError{
				// 	Success: false,
				// 	Data:    nil,
				// 	Message: fiber.ErrInternalServerError.Message,
				// 	Error:   nil,
				// }
				// return ctx.Status(fiber.StatusInternalServerError).JSON(res)
			}

			return nil
		},
	})

	app.Use(recoverFiber.New())
	app.Use(ErrorHandler())
	// CORS Config
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PATCH,DELETE,PUT",
	}))

	app.Use(logger.New())

	api := app.Group("/api/v1")
	routes.RootRoute(api)
	app.Listen(port)
}
