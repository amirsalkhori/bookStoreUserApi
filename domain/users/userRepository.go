package users

import (
	"bookStoreUser/errors"
	dateutils "bookStoreUser/utils/dateUtils"
	"fmt"
)

var (
	userDb = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	result := userDb[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not fount", user.Id))
	}
	user.Id = result.Id
	user.Name = result.Name
	user.Family = result.Family
	user.Email = result.Email
	user.CreatedAt = result.CreatedAt

	return nil
}

func (user *User) Save() *errors.RestError {
	currentUser := userDb[user.Id]
	if currentUser != nil {
		if currentUser.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s is already exists!", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("Id %d is already exists!", user.Id))
	}
	userDb[user.Id] = user
	user.CreatedAt = dateutils.GetNowString()

	return nil
}
