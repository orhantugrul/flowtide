package user

import (
	"time"

	"gorm.io/gorm"
)

type UserSchema struct {
	ID        uint           `json:"id"`
	Username  string         `json:"username"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type UserCreateSchema struct {
	Username string `validate:"required"`
}

func (user *UserCreateSchema) ToModel() User {
	return User{
		Username: user.Username,
	}
}

type UserUpdateSchema struct {
	Username string `validate:"required"`
}

func (user *UserUpdateSchema) ToModel() User {
	return User{
		Username: user.Username,
	}
}
