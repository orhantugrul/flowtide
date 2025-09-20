package main

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/orhantugrul/flowtide/app/activity"
	"github.com/orhantugrul/flowtide/app/editor"
	"github.com/orhantugrul/flowtide/app/project"
	"github.com/orhantugrul/flowtide/app/user"
	"github.com/orhantugrul/flowtide/database"
)

func main() {
	app := fiber.New(fiber.Config{ErrorHandler: errorHandler})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	defer func() {
		if err := database.Close(); err != nil {
			log.Error("Failed to close database: ", err)
		}
	}()

	router := app.Group("/api")
	{
		activity.UseRoutes(router)
		editor.UseRoutes(router)
		project.UseRoutes(router)
		user.UseRoutes(router)
	}

	log.Fatal(app.Listen(":8080"))
}

func errorHandler(context *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	message := err.Error()

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		code = fiberErr.Code
		message = fiberErr.Message
	}

	errorResponse := map[string]any{
		"code":      code,
		"message":   message,
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"path":      context.Path(),
	}

	return context.Status(code).JSON(errorResponse)
}
