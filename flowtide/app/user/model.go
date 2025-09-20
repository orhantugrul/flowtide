package user

import (
	"github.com/orhantugrul/flowtide/app/project"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Projects []project.Project
}

func (user *User) ToSchema() UserSchema {
	return UserSchema{
		ID:       user.ID,
		Username: user.Username,
	}
}
