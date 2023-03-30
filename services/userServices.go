package services

import (
	"bookStoreUser/domain/users"
	"bookStoreUser/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return &user, nil
}
