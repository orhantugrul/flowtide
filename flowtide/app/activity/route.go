package activity

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate = validator.New()

func UseRoutes(router fiber.Router) {
	activities := router.Group("/activities")
	activities.Get("/", getActivities)
	activities.Get("/:id", getActivity)
	activities.Post("/", createActivity)
	activities.Put("/:id", updateActivity)
	activities.Delete("/:id", deleteActivity)
}

func getActivities(context *fiber.Ctx) error {
	activities, err := GetActivities()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	schemas := []ActivitySchema{}
	for _, activity := range activities {
		schemas = append(schemas, activity.ToSchema())
	}

	return context.JSON(schemas)
}

func getActivity(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "provide a valid id")
	}

	activity, err := GetActivity(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return context.JSON(activity.ToSchema())
}

func createActivity(context *fiber.Ctx) error {
	schema := ActivityCreateSchema{}
	if err := context.BodyParser(&schema); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(schema); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	activity, err := CreateActivity(&schema)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	location := fmt.Sprintf("%s/%d", context.Path(), activity.ID)
	context.Set("Location", location)
	context.Status(fiber.StatusCreated)
	return nil
}

func updateActivity(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "provide a valid id")
	}

	schema := ActivityUpdateSchema{}
	if err := context.BodyParser(&schema); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(schema); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	activity, err := UpdateActivity(id, &schema)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	location := fmt.Sprintf("%s/%d", context.Path(), activity.ID)
	context.Set("Location", location)
	context.Status(fiber.StatusOK)
	return nil
}

func deleteActivity(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "provide a valid id")
	}

	if err := DeleteActivity(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	context.Status(fiber.StatusNoContent)
	return nil
}
