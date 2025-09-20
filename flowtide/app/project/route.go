package project

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/orhantugrul/flowtide/app/activity"
)

var validate *validator.Validate = validator.New()

func UseRoutes(router fiber.Router) {
	projects := router.Group("/projects")
	projects.Get("/", getProjects)
	projects.Get("/:id", getProject)
	projects.Get("/:id/activities", getProjectActivities)
	projects.Post("/", createProject)
	projects.Put("/:id", updateProject)
	projects.Delete("/:id", deleteProject)
}

func getProjects(context *fiber.Ctx) error {
	projects, err := GetProjects()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	schemas := []ProjectSchema{}
	for _, project := range projects {
		schemas = append(schemas, project.ToSchema())
	}

	return context.JSON(schemas)
}

func getProject(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "provide a valid id")
	}

	project, err := GetProject(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return context.JSON(project.ToSchema())
}

func getProjectActivities(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "provide a valid id")
	}

	activities, err := GetProjectActivities(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	schemas := []activity.ActivitySchema{}
	for _, activity := range activities {
		schemas = append(schemas, activity.ToSchema())
	}

	return context.JSON(schemas)
}

func createProject(context *fiber.Ctx) error {
	schema := ProjectCreateSchema{}
	if err := context.BodyParser(&schema); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(schema); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	project, err := CreateProject(&schema)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	location := fmt.Sprintf("%s/%d", context.Path(), project.ID)
	context.Set("Location", location)
	context.Status(fiber.StatusCreated)
	return context.JSON(project.ToSchema())
}

func updateProject(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "provide a valid id")
	}

	schema := ProjectUpdateSchema{}
	if err := context.BodyParser(&schema); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(schema); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	project, err := UpdateProject(id, &schema)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	location := fmt.Sprintf("%s/%d", context.Path(), project.ID)
	context.Set("Location", location)
	return context.JSON(project.ToSchema())
}

func deleteProject(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "provide a valid id")
	}

	if err := DeleteProject(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	context.Status(fiber.StatusNoContent)
	return nil
}
