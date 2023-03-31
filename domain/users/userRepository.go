package users

import (
	usersDB "bookStoreUser/dataSorces/mysql/usersDB"
	"bookStoreUser/errors"
	"bookStoreUser/utils/dateUtils"
	"fmt"
	"strings"
)

var (
	userDb = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	if err := usersDB.Client.Ping(); err != nil {
		panic(err)
	}
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
	query := "INSERT INTO users (name, family,email, created_at) VALUES (?, ?, ?, ?);"
	statement, err := usersDB.Client.Prepare(query)
	if err != nil {
		return errors.NewInternamlServerError(err.Error())
	}
	defer statement.Close()
	user.CreatedAt = dateUtils.GetNowString()

	fmt.Println("", user.CreatedAt)
	insertResult, err := statement.Exec(user.Name, user.Family, user.Email, user.CreatedAt)
	if err != nil {
		if strings.Contains(err.Error(), "email_UNIQUE") {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already exsist!", user.Email))
		}
		return errors.NewInternamlServerError("Error durring the insert user")
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternamlServerError("Error durring save user")
	}
	user.Id = userId

	return nil
}
