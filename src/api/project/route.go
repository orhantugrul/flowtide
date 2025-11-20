package project

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/orhantugrul/flowtide/src/api/activity"
	"github.com/orhantugrul/flowtide/src/validator"
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
	query := ProjectQueryInput{}
	if err := context.QueryParser(&query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	projects, err := GetProjects(&query)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	data := []ProjectOutput{}
	for _, item := range projects {
		data = append(data, ProjectOutput{
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
	params := ProjectParamsInput{}
	if err := context.ParamsParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	project, err := GetProject(params.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return context.JSON(ProjectOutput{
		ID:        project.ID,
		Name:      project.Name,
		Path:      project.Path,
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
		DeletedAt: project.DeletedAt,
	})
}

func getProjectActivities(context *fiber.Ctx) error {
	params := ProjectParamsInput{}
	if err := context.ParamsParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	activities, err := GetProjectActivities(params.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	data := []activity.ActivityOutput{}
	for _, item := range activities {
		data = append(data, activity.ActivityOutput{
			ID:        item.ID,
			ProjectID: item.ProjectID,
			EditorID:  item.EditorID,
			Language:  item.Language,
			FilePath:  item.FilePath,
			StartedAt: item.StartedAt,
			EndedAt:   item.EndedAt,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
			DeletedAt: item.DeletedAt,
		})
	}

	return context.JSON(data)
}

func createProject(context *fiber.Ctx) error {
	body := ProjectCreateInput{}
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
	return context.JSON(ProjectOutput{
		ID:        project.ID,
		Name:      project.Name,
		Path:      project.Path,
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
		DeletedAt: project.DeletedAt,
	})
}

func updateProject(context *fiber.Ctx) error {
	params := ProjectParamsInput{}
	if err := context.ParamsParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body := ProjectUpdateInput{}
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
	return context.JSON(ProjectOutput{
		ID:        project.ID,
		Name:      project.Name,
		Path:      project.Path,
		CreatedAt: project.CreatedAt,
		UpdatedAt: project.UpdatedAt,
		DeletedAt: project.DeletedAt,
	})
}

func deleteProject(context *fiber.Ctx) error {
	params := ProjectParamsInput{}
	if err := context.ParamsParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := DeleteProject(params.ID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	context.Status(fiber.StatusNoContent)
	return nil
}
