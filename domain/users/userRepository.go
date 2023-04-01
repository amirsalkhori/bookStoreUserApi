package users

import (
	usersDB "bookStoreUser/dataSorces/mysql/usersDB"
	"bookStoreUser/errors"
	"bookStoreUser/utils/dateUtils"
	"fmt"
	"strings"
)


func (user *User) Get() *errors.RestError {
	if err := usersDB.Client.Ping(); err != nil {
		panic(err)
	}
	query := "SELECT * FROM users where id = ?"
	statement, err := usersDB.Client.Prepare(query)
	if err != nil {
		return errors.NewInternamlServerError(err.Error())
	}
	defer statement.Close()
	result := statement.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.Name, &user.Family, &user.Email, &user.CreatedAt); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
		}
		return errors.NewInternamlServerError(fmt.Sprintf("Error when trying get user %d: %s", user.Id, err.Error()))
	}

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

func (user *User) Update() *errors.RestError {
	query := "UPDATE users SET name = ?, family = ?,email = ? where id = ?;"
	statement, err := usersDB.Client.Prepare(query)
	if err != nil {
		return errors.NewInternamlServerError(err.Error())
	}
	defer statement.Close()
	
	res, err := statement.Exec(user.Name, user.Family, user.Email, user.Id)
	fmt.Println(res)
	if err != nil {
		if strings.Contains(err.Error(), "email_UNIQUE") {
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already exsist!", user.Email))
		}
		return errors.NewInternamlServerError("Error durring the update user")
	}
	
	return nil
}