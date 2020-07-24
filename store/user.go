package store

import (
	"errors"
	"user/model"

	"github.com/jinzhu/gorm"
)

var ErrNotFound = errors.New("this user doesn't exist in the database")
var ErrWrongPass = errors.New("password is not correct")

type User interface {
	Insert(user model.User) error
	Retrieve(user model.User) (model.User, error)
}

type SQLUser struct {
	DB *gorm.DB
}

func NewUser(d *gorm.DB) SQLUser {
	return SQLUser{DB: d}
}

func (u SQLUser) Insert(user model.User) error {
	result := u.DB.Create(&user)

	return result.Error
}

func (u SQLUser) Retrieve(user model.User) (model.User, error) {
	var us model.User

	u.DB.Where("email = ?", user.Email).First(&us)

	var err error

	if us.Email == "" {
		err = ErrNotFound
		return us, err
	}

	if us.Password != user.Password {
		err = ErrWrongPass
		return us, err
	}

	return us, err
}
