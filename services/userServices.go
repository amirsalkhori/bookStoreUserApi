package services

import (
	"bookStoreUser/domain/users"
	"bookStoreUser/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateUser(user users.User) (*users.User, *errors.RestError) {
	currentUser, errResult := GetUser(user.Id)
	if errResult != nil {
		return nil, errResult
	}

	if user.Name != "" {
		currentUser.Name = user.Name
	}
	if user.Family != "" {
		currentUser.Family = user.Family
	}
	if user.Email != "" {
		currentUser.Email = user.Email
	}

	if err := currentUser.Update(); err != nil {
		return nil, err
	}
	return currentUser, nil
}
