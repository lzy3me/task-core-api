package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"task-core/routes"
	"time"

	db "task-core/db"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	recoverFiber "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"server":    "Server is up and running",
		"db_status": "Connected",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if errP := db.MG.Client.Ping(ctx, readpref.Primary()); errP != nil {
		log.Println("errP", errP)
		res["db_status"] = "Not Connected"
	}

	return c.JSON(res)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error while loading env file.")
	}

	port := os.Getenv("PORT")

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

	db.Connect()
	app.Use(logger.New())
	app.Get("/healthcheck", HealthCheck)

	api := app.Group("/api/v1")
	routes.RootRoute(api)
	app.Listen(port)
}
