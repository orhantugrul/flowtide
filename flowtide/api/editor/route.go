package editor

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/orhantugrul/flowtide/validator"
)

func UseRoutes(router fiber.Router) {
	editors := router.Group("/editors")
	editors.Get("/", getEditors)
	editors.Get("/:id", getEditor)
	editors.Post("/", createEditor)
	editors.Put("/:id", updateEditor)
	editors.Delete("/:id", deleteEditor)
}

func getEditors(context *fiber.Ctx) error {
	query := EditorQueryInput{}
	if err := context.QueryParser(&query); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	editors, err := GetEditors(&query)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	data := []EditorOutput{}
	for _, editor := range editors {
		data = append(data, EditorOutput{
			ID:        editor.ID,
			Name:      editor.Name,
			Version:   editor.Version,
			CreatedAt: editor.CreatedAt,
			UpdatedAt: editor.UpdatedAt,
			DeletedAt: editor.DeletedAt,
		})
	}

	return context.JSON(data)
}

func getEditor(context *fiber.Ctx) error {
	params := EditorParamsInput{}
	if err := context.ParamsParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	editor, err := GetEditor(params.ID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return context.JSON(EditorOutput{
		ID:        editor.ID,
		Name:      editor.Name,
		Version:   editor.Version,
		CreatedAt: editor.CreatedAt,
		UpdatedAt: editor.UpdatedAt,
		DeletedAt: editor.DeletedAt,
	})
}

func createEditor(context *fiber.Ctx) error {
	body := EditorCreateInput{}
	if err := context.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate(&body); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	editor, err := CreateEditor(&body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	location := fmt.Sprintf("%s/%d", context.Path(), editor.ID)
	context.Set("Location", location)
	context.Status(fiber.StatusCreated)
	return context.JSON(EditorOutput{
		ID:        editor.ID,
		Name:      editor.Name,
		Version:   editor.Version,
		CreatedAt: editor.CreatedAt,
		UpdatedAt: editor.UpdatedAt,
		DeletedAt: editor.DeletedAt,
	})
}

func updateEditor(context *fiber.Ctx) error {
	params := EditorParamsInput{}
	if err := context.ParamsParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	body := EditorUpdateInput{}
	if err := context.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validator.Validate(&body); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	editor, err := UpdateEditor(params.ID, &body)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	location := fmt.Sprintf("%s/%d", context.Path(), editor.ID)
	context.Set("Location", location)
	return context.JSON(EditorOutput{
		ID:        editor.ID,
		Name:      editor.Name,
		Version:   editor.Version,
		CreatedAt: editor.CreatedAt,
		UpdatedAt: editor.UpdatedAt,
		DeletedAt: editor.DeletedAt,
	})
}

func deleteEditor(context *fiber.Ctx) error {
	params := EditorParamsInput{}
	if err := context.ParamsParser(&params); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := DeleteEditor(params.ID); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	context.Status(fiber.StatusNoContent)
	return nil
}
