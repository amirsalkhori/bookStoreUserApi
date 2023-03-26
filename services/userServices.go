package services

import (
	"bookStoreUser/domain/users"
	"bookStoreUser/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	return &user, nil
}
