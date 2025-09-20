package user

import (
	"context"

	"github.com/orhantugrul/flowtide/app/project"
	"github.com/orhantugrul/flowtide/database"
	"gorm.io/gorm"
)

func GetUsers() ([]User, error) {
	return gorm.G[User](database.Database).Find(context.Background())
}

func GetUser(id int) (User, error) {
	return gorm.G[User](database.Database).
		Where("id = ?", id).
		First(context.Background())
}

func GetUserProjects(id int) ([]project.Project, error) {
	return gorm.G[project.Project](database.Database).
		Where("user_id = ?", id).
		Find(context.Background())
}

func CreateUser(schema *UserCreateSchema) (*User, error) {
	user := schema.ToModel()

	err := gorm.G[User](database.Database).
		Create(context.Background(), &user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(id int, schema *UserUpdateSchema) (*User, error) {
	user := schema.ToModel()

	_, err := gorm.G[User](database.Database).
		Where("id = ?", id).
		Updates(context.Background(), user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func DeleteUser(id int) error {
	_, err := gorm.G[User](database.Database).
		Where("id = ?", id).
		Delete(context.Background())

	if err != nil {
		return err
	}
	return nil
}
