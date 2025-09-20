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
}

func getActivities(context *fiber.Ctx) error {
	activities, err := GetActivities()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	schemas := []ActivitySchema{}
	for _, activity := range activities {
		schemas = append(schemas, activity.toSchema())
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

	schema := activity.toSchema()
	return context.JSON(schema)
}

func createActivity(context *fiber.Ctx) error {
	schema := ActivityCreateSchema{}
	if err := context.BodyParser(&schema); err != nil {
		fiber.NewError(fiber.StatusBadRequest, err.Error())
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
