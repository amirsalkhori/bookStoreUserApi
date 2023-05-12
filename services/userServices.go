package services

import (
	"bookStoreUser/domain/users"
	"bookStoreUser/errors"
	"bookStoreUser/utils/cryptoUtils"

	"github.com/gin-gonic/gin"
)

var(
	UserService userServiceInterface = &userService{}
)

type userService struct{

}

type userServiceInterface interface{
	CreateUser(users.User) (*users.User, *errors.RestError)
	GetUser(int64) (*users.User, *errors.RestError)
	UpdateUser(users.User) (*users.User, *errors.RestError)
	DeleteUser(users.User) *errors.RestError
	GetUserCollection(*gin.Context) (users.Users, *errors.RestError)
	GetUserByStatus(bool) (users.Users, *errors.RestError)
}

func(userService *userService) CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Password = cryptoUtils.GetMd5(user.Password)
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func(userService *userService) GetUser(userId int64) (*users.User, *errors.RestError) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func(userService *userService) UpdateUser(user users.User) (*users.User, *errors.RestError) {
	currentUser, errResult := userService.GetUser(user.Id)
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

func(userService *userService) DeleteUser(user users.User) *errors.RestError {
	currentUser, errResult := userService.GetUser(user.Id)
	if errResult != nil {
		return errResult
	}

	if err := currentUser.Delete(); err != nil {
		return err
	}

	return nil
}

func(userService *userService) GetUserCollection(c *gin.Context) (users.Users, *errors.RestError) {
	result := users.User{}
	users, err := result.GetCollection()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func(userService *userService) GetUserByStatus(status bool) (users.Users, *errors.RestError) {
	users := &users.User{}
	result, err := users.FibdByStatus(status)
	if err != nil {
		return nil, err
	}

	return result, nil
}
