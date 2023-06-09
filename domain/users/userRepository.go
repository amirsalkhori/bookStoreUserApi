package users

import (
	usersDB "bookStoreUser/dataSorces/mysql/usersDB"
	"bookStoreUser/errors"
	"bookStoreUser/logger"
	"bookStoreUser/utils/dateUtils"
	"bookStoreUser/utils/elastic"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

func (user *User) Get() *errors.RestError {
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

func connection() (*gorm.DB, *errors.RestError) {
	db, err := usersDB.Connect()
	if err != nil {
		logger.Error("Error when tryin to connect to db", err)
		return nil, errors.NewInternamlServerError(err.Error())
	}

	return db, nil
}

func (user *User) Save() *errors.RestError {
	db, errorConnection := connection()
	if errorConnection != nil {
		return errorConnection
	}
	user.CreatedAt = dateUtils.GetNowString()
	userItem := User{Name: user.Name, Family: user.Family, Email: user.Email, Status: user.Status, Password: user.Password, CreatedAt: user.CreatedAt}

	err := db.Create(&userItem).Error
	if err != nil {
		if strings.Contains(err.Error(), "email_UNIQUE") {
			logger.Error("Error when tryin to save user, duplicate error", err)
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already exsist!", user.Email))
		}
		logger.Error("Error when tryin to save user", err)
		return errors.NewInternamlServerError("Error durring the insert user")
	}
	user.Id = userItem.Id

	elasticErr := elastic.Elastic.Insert(elastic.Elastic{}, "user_index", user)
	if elasticErr != nil {
		logger.Error("Error during index user into elastic", err)
		return nil
	}

	return nil
}

func (user *User) Update() *errors.RestError {
	db, errorConnection := connection()
	if errorConnection != nil {
		return errorConnection
	}
	err := db.Save(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "email_UNIQUE") {
			logger.Error("Error when tryin to update user, duplicate error", err)
			return errors.NewBadRequestError(fmt.Sprintf("Email %s already exsist!", user.Email))
		}
	}

	return nil
}

func (user *User) Delete() *errors.RestError {
	db, errorConnection := connection()
	if errorConnection != nil {
		return errorConnection
	}
	err := db.Delete(&user).Error
	if err != nil {
		logger.Error("Error when tryin to delete user", err)
		return errors.NewInternamlServerError(err.Error())
	}

	return nil
}

func (user *User) GetCollection() ([]User, *errors.RestError) {
	users := make([]User, 0)
	db, errorConnection := connection()
	if errorConnection != nil {
		return nil, errorConnection
	}
	db.Find(&users) // SELECT * FROM users;

	return users, nil
}

func (user User) FibdByStatus(status bool) ([]User, *errors.RestError) {
	users := make([]User, 0)
	db, errorConnection := connection()
	if errorConnection != nil {
		return nil, errorConnection
	}
	err := db.Where("status = ?", status).Find(&users).Error
	if err != nil {
		logger.Error("Error when tryin to get users", err)
		return nil, errors.NewInternamlServerError(err.Error())
	}

	return users, nil
}

func (u *User) GetUserByEmailAndPassword() (*User, *errors.RestError) {
	user := User{}
	db, errorConnection := connection()
	if errorConnection != nil {
		return nil, errorConnection
	}
	err := db.Where("email = ? AND password = ? AND status = 1", u.Email, u.Password).Find(&user).Error
	if err != nil {
		logger.Error("Error when tryin to get users", err)
		return nil, errors.NewInternamlServerError(err.Error())
	}
	if user.Id == 0 {
		return nil, errors.NewNotFoundError("Invalid credential")
	}

	return &user, nil
}
