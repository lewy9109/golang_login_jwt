package user

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string
	LastName    string
	Email       string
	Password    string
	AccessToken string
}
