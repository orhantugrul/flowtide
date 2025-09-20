package editor

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate = validator.New()

func UseRoutes(router fiber.Router) {
	editors := router.Group("/editors")
	editors.Get("/", getEditors)
	editors.Get("/:id", getEditor)
	editors.Post("/", createEditor)
	editors.Put("/:id", updateEditor)
	editors.Delete("/:id", deleteEditor)
}

func getEditors(context *fiber.Ctx) error {
	editors, err := GetEditors()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	schemas := []EditorSchema{}
	for _, editor := range editors {
		schemas = append(schemas, editor.ToSchema())
	}

	return context.JSON(schemas)
}

func getEditor(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "provide a valid id")
	}

	editor, err := GetEditor(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return context.JSON(editor.ToSchema())
}

func createEditor(context *fiber.Ctx) error {
	schema := EditorCreateSchema{}
	if err := context.BodyParser(&schema); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(schema); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	editor, err := CreateEditor(&schema)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	location := fmt.Sprintf("%s/%d", context.Path(), editor.ID)
	context.Set("Location", location)
	context.Status(fiber.StatusCreated)
	return context.JSON(editor.ToSchema())
}

func updateEditor(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "provide a valid id")
	}

	schema := EditorUpdateSchema{}
	if err := context.BodyParser(&schema); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(schema); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	editor, err := UpdateEditor(id, &schema)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	location := fmt.Sprintf("%s/%d", context.Path(), editor.ID)
	context.Set("Location", location)
	return context.JSON(editor.ToSchema())
}

func deleteEditor(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "provide a valid id")
	}

	if err := DeleteEditor(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	context.Status(fiber.StatusNoContent)
	return nil
}
