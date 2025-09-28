package project

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/orhantugrul/flowtide/api/activity"
	"github.com/orhantugrul/flowtide/validator"
)

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

	data := []ProjectSchema{}
	for _, item := range projects {
		data = append(data, ProjectSchema{
			ID:        item.ID,
			Name:      item.Name,
			Path:      item.Path,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
			DeletedAt: item.DeletedAt,
		})
	}

	return context.JSON(data)
}

func getProject(context *fiber.Ctx) error {
	params := ProjectParamsSchema{}
	if err := context.ParamsParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	project, err := GetProject(params.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return context.JSON(ProjectSchema{
		ID:        project.ID,
		Name:      project.Name,
		Path:      project.Path,
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
		DeletedAt: project.DeletedAt,
	})
}

func getProjectActivities(context *fiber.Ctx) error {
	params := ProjectParamsSchema{}
	if err := context.ParamsParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	activities, err := GetProjectActivities(params.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	data := []activity.ActivitySchema{}
	for _, item := range activities {
		data = append(data, activity.ActivitySchema{
			ID:        item.ID,
			ProjectID: item.ProjectID,
			EditorID:  item.EditorID,
			Language:  item.Language,
			FilePath:  item.FilePath,
			StartTime: item.StartTime,
			EndTime:   item.EndTime,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
			DeletedAt: item.DeletedAt,
		})
	}

	return context.JSON(data)
}

func createProject(context *fiber.Ctx) error {
	body := ProjectCreateSchema{}
	if err := context.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate(&body); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	project, err := CreateProject(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	location := fmt.Sprintf("%s/%d", context.Path(), project.ID)
	context.Set("Location", location)
	context.Status(fiber.StatusCreated)
	return context.JSON(ProjectSchema{
		ID:        project.ID,
		Name:      project.Name,
		Path:      project.Path,
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
		DeletedAt: project.DeletedAt,
	})
}

func updateProject(context *fiber.Ctx) error {
	params := ProjectParamsSchema{}
	if err := context.ParamsParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body := ProjectUpdateSchema{}
	if err := context.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate(&body); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	project, err := UpdateProject(params.ID, &body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	location := fmt.Sprintf("%s/%d", context.Path(), project.ID)
	context.Set("Location", location)
	return context.JSON(ProjectSchema{
		ID:        project.ID,
		Name:      project.Name,
		Path:      project.Path,
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
		DeletedAt: project.DeletedAt,
	})
}

func deleteProject(context *fiber.Ctx) error {
	params := ProjectParamsSchema{}
	if err := context.ParamsParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := DeleteProject(params.ID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	context.Status(fiber.StatusNoContent)
	return nil
}
