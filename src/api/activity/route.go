package activity

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/orhantugrul/flowtide/src/validator"
)

func UseRoutes(router fiber.Router) {
	activities := router.Group("/activities")
	activities.Get("/", getActivities)
	activities.Get("/:id", getActivity)
	activities.Post("/", createActivity)
	activities.Post("/batch", createActivities)
	activities.Put("/:id", updateActivity)
	activities.Delete("/:id", deleteActivity)
}

func getActivities(context *fiber.Ctx) error {
	query := ActivityQueryInput{}
	if err := context.QueryParser(&query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	activities, err := GetActivities(&query)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	data := []ActivityOutput{}
	for _, activity := range activities {
		data = append(data, ActivityOutput{
			ID:        activity.ID,
			ProjectID: activity.ProjectID,
			EditorID:  activity.EditorID,
			Language:  activity.Language,
			FilePath:  activity.FilePath,
			StartedAt: activity.StartedAt,
			EndedAt:   activity.EndedAt,
			CreatedAt: activity.CreatedAt,
			UpdatedAt: activity.UpdatedAt,
			DeletedAt: activity.DeletedAt,
		})
	}

	return context.JSON(data)
}

func getActivity(context *fiber.Ctx) error {
	params := ActivityParamsInput{}
	if err := context.ParamsParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	activity, err := GetActivity(params.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return context.JSON(ActivityOutput{
		ID:        activity.ID,
		ProjectID: activity.ProjectID,
		EditorID:  activity.EditorID,
		Language:  activity.Language,
		FilePath:  activity.FilePath,
		StartedAt: activity.StartedAt,
		EndedAt:   activity.EndedAt,
		CreatedAt: activity.CreatedAt,
		UpdatedAt: activity.UpdatedAt,
		DeletedAt: activity.DeletedAt,
	})
}

func createActivity(context *fiber.Ctx) error {
	body := ActivityCreateInput{}
	if err := context.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate(&body); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	activity, err := CreateActivity(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	location := fmt.Sprintf("%s/%d", context.Path(), activity.ID)
	context.Set("Location", location)
	context.Status(fiber.StatusCreated)
	return context.JSON(ActivityOutput{
		ID:        activity.ID,
		ProjectID: activity.ProjectID,
		EditorID:  activity.EditorID,
		Language:  activity.Language,
		FilePath:  activity.FilePath,
		StartedAt: activity.StartedAt,
		EndedAt:   activity.EndedAt,
		CreatedAt: activity.CreatedAt,
		UpdatedAt: activity.UpdatedAt,
		DeletedAt: activity.DeletedAt,
	})
}

func createActivities(context *fiber.Ctx) error {
	body := []ActivityCreateInput{}
	if err := context.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	for _, item := range body {
		if err := validator.Validate(&item); err != nil {
			return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
		}
	}

	activities, err := CreateActivities(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return context.JSON(activities)
}

func updateActivity(context *fiber.Ctx) error {
	params := ActivityParamsInput{}
	if err := context.ParamsParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body := ActivityUpdateInput{}
	if err := context.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate(&body); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	activity, err := UpdateActivity(params.ID, &body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	location := fmt.Sprintf("%s/%d", context.Path(), activity.ID)
	context.Set("Location", location)
	return context.JSON(ActivityOutput{
		ID:        activity.ID,
		ProjectID: activity.ProjectID,
		EditorID:  activity.EditorID,
		Language:  activity.Language,
		FilePath:  activity.FilePath,
		StartedAt: activity.StartedAt,
		EndedAt:   activity.EndedAt,
		CreatedAt: activity.CreatedAt,
		UpdatedAt: activity.UpdatedAt,
		DeletedAt: activity.DeletedAt,
	})
}

func deleteActivity(context *fiber.Ctx) error {
	params := ActivityParamsInput{}
	if err := context.ParamsParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := DeleteActivity(params.ID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	context.Status(fiber.StatusNoContent)
	return nil
}
