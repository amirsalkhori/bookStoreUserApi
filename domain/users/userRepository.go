package users

import (
	usersDB "bookStoreUser/dataSorces/mysql/usersDB"
	"bookStoreUser/errors"
	"bookStoreUser/logger"
	"bookStoreUser/utils/dateUtils"
	"fmt"
	"strings"
)

func (user *User) Get() *errors.RestError {
	// if err := usersDB.Client.Ping(); err != nil {
	// 	panic(err)
	// }
	// query := "SELECT * FROM users where id = ?"
	// statement, err := usersDB.Client.Prepare(query)
	// if err != nil {
	// 	return errors.NewInternamlServerError(err.Error())
	// }
	// defer statement.Close()
	// result := statement.QueryRow(user.Id)
	// if err := result.Scan(&user.Id, &user.Name, &user.Family, &user.Email, &user.Status, &user.Password, &user.CreatedAt); err != nil {
	// 	if strings.Contains(err.Error(), "no rows in result set") {
	// 		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	// 	}
	// 	return errors.NewInternamlServerError(fmt.Sprintf("Error when trying get user %d: %s", user.Id, err.Error()))
	// }

	db, err := usersDB.Connect()
	if err != nil {
		logger.Error("Error when tryin to connect to db", err)
		return errors.NewInternamlServerError(err.Error())
	}
	err = db.Where("id = ?", user.Id).Find(&user).Error
	if err != nil {
		logger.Error("Error when tryin to get user", err)
		return errors.NewInternamlServerError(err.Error())
	}

	return nil
}

func (user *User) Save() *errors.RestError {
	// query := "INSERT INTO users (name, family,email, status, password, created_at) VALUES (?, ?, ?, ?, ?, ?);"
	// statement, err := usersDB.Client.Prepare(query)
	// if err != nil {
	// 	return errors.NewInternamlServerError(err.Error())
	// }
	// defer statement.Close()
	// user.CreatedAt = dateUtils.GetNowString()

	// insertResult, err := statement.Exec(user.Name, user.Family, user.Email, user.Status, user.Password, user.CreatedAt)
	// if err != nil {
	// 	if strings.Contains(err.Error(), "email_UNIQUE") {
	// 		return errors.NewBadRequestError(fmt.Sprintf("Email %s already exsist!", user.Email))
	// 	}
	// 	return errors.NewInternamlServerError("Error durring the insert user")
	// }
	// userId, err := insertResult.LastInsertId()
	// if err != nil {
	// 	return errors.NewInternamlServerError("Error durring save user")
	// }
	// user.Id = userId
	db, err := usersDB.Connect()
	if err != nil {
		logger.Error("Error when tryin to connect to db", err)
		return errors.NewInternamlServerError(err.Error())
	}
	user.CreatedAt = dateUtils.GetNowString()
	userItem := User{Name: user.Name, Family: user.Family, Email: user.Email, Status: user.Status, Password: user.Password, CreatedAt: user.CreatedAt}

	err = db.Create(&userItem).Error
	if err != nil {
		if strings.Contains(err.Error(), "email_UNIQUE") {
			logger.Error("Error when tryin to save user, duplicate error", err)
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already exsist!", user.Email))
		}
		logger.Error("Error when tryin to save user", err)
		return errors.NewInternamlServerError("Error durring the insert user")
	}
	user.Id = userItem.Id

	return nil
}

func (user *User) Update() *errors.RestError {
	db, err := usersDB.Connect()
	if err != nil {
		logger.Error("Error when tryin to connect to db", err)
		return errors.NewInternamlServerError(err.Error())
	}
	err = db.Save(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "email_UNIQUE") {
			logger.Error("Error when tryin to update user, duplicate error", err)
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already exsist!", user.Email))
		}
	}

	return nil
}

func (user *User) Delete() *errors.RestError {
	db, err := usersDB.Connect()
	if err != nil {
		logger.Error("Error when tryin to connect to db", err)
		return errors.NewInternamlServerError(err.Error())
	}
	err = db.Delete(&user).Error
	if err != nil {
		logger.Error("Error when tryin to delete user", err)
		return errors.NewInternamlServerError(err.Error())
	}

	return nil
}

func (user *User) GetCollection() ([]User, *errors.RestError) {
	users := make([]User, 0)
	db, err := usersDB.Connect()
	if err != nil {
		logger.Error("Error when tryin to connect to db", err)
		return nil, errors.NewInternamlServerError(err.Error())
	}
	db.Find(&users) // SELECT * FROM users;

	return users, nil
}

func (user User) FibdByStatus(status bool) ([]User, *errors.RestError) {
	users := make([]User, 0)
	db, err := usersDB.Connect()
	if err != nil {
		logger.Error("Error when tryin to connect to db", err)
		return nil, errors.NewInternamlServerError(err.Error())
	}
	err = db.Where("status = ?", status).Find(&users).Error
	if err != nil {
		logger.Error("Error when tryin to get users", err)
		return nil, errors.NewInternamlServerError(err.Error())
	}
	db.Where("status = ?", status).
		Find(&users) // SELECT * FROM users;

	return users, nil
}
