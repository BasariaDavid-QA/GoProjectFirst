package userService

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"Email"`
	Password string `json:"Password"`
}
