package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
}

func (user *User) ToSchema() UserSchema {
	return UserSchema{
		ID:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}
