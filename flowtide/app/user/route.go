package user

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/orhantugrul/flowtide/app/project"
)

func UseRoutes(router fiber.Router) {
	users := router.Group("/users")
	users.Get("/", getUsers)
	users.Get("/:id", getUser)
	users.Get("/:id/projects", getUserProjects)
	users.Post("/", createUser)
	users.Put("/:id", updateUser)
	users.Delete("/:id", deleteUser)
}

var validate *validator.Validate = validator.New()

func getUsers(context *fiber.Ctx) error {
	users, err := GetUsers()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	schemas := []UserSchema{}
	for _, user := range users {
		schemas = append(schemas, user.ToSchema())
	}

	return context.JSON(schemas)
}

func getUser(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "provide a valid id")
	}

	user, err := GetUser(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return context.JSON(user.ToSchema())
}

func getUserProjects(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "provide a valid id")
	}

	projects, err := GetUserProjects(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	schemas := []project.ProjectSchema{}
	for _, project := range projects {
		schemas = append(schemas, project.ToSchema())
	}

	return context.JSON(schemas)
}

func createUser(context *fiber.Ctx) error {
	schema := UserCreateSchema{}
	if err := context.BodyParser(&schema); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(schema); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	user, err := CreateUser(&schema)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	location := fmt.Sprintf("%s/%d", context.Path(), user.ID)
	context.Set("Location", location)
	context.Status(fiber.StatusCreated)
	return context.JSON(user.ToSchema())
}

func updateUser(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "provide a valid id")
	}

	schema := UserUpdateSchema{}
	if err := context.BodyParser(&schema); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if err := validate.Struct(schema); err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	user, err := UpdateUser(id, &schema)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	location := fmt.Sprintf("%s/%d", context.Path(), user.ID)
	context.Set("Location", location)
	return context.JSON(user.ToSchema())
}

func deleteUser(context *fiber.Ctx) error {
	id, err := context.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "provide a valid id")
	}

	if err := DeleteUser(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	context.Status(fiber.StatusNoContent)
	return nil
}
