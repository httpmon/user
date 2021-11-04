package mock

import (
	"errors"

	"github.com/httpmon/user/model"
)

var (
	ErrDuplicateEmail = errors.New("this email exists")
	ErrWrongPass      = errors.New("password is not correct")
	ErrNotFound       = errors.New("this github.com/httpmon/user doesn't exist in the database")
)

type User struct {
	Info map[string]string
}

func (u User) Insert(user model.User) error {
	if _, ok := u.Info[user.Email]; ok {
		return ErrDuplicateEmail
	}

	u.Info[user.Email] = user.Password

	return nil
}

func (u User) Retrieve(user model.User) (model.User, error) {
	pass, ok := u.Info[user.Email]
	if ok {
		if user.Password == pass {
			return model.User{
				ID:       0,
				Email:    user.Email,
				Password: pass,
				Urls:     nil,
			}, nil
		}

		return model.User{}, ErrWrongPass
	}

	return model.User{}, ErrNotFound
}
