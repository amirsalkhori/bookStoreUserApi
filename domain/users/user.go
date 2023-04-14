package users

import (
	"bookStoreUser/errors"
	"strings"
)

type User struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	Family    string `json:"family"`
	Email     string `json:"email"`
	CreatedAt string `json:"createdAt"`
	Status    bool   `json:"status"`
	Password  string `json:"password"`
}

func (user *User) Validate() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address!")
	}
	user.Password = strings.TrimSpace(strings.ToLower(user.Password))
	if user.Password == "" {
		return errors.NewBadRequestError("Password is requered!")
	}

	return nil
}
