package controllers

import (
	"bookStoreUser/domain/users"
	"bookStoreUser/errors"
	"bookStoreUser/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	// bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//TODO: Handle error
	// 	return
	// }
	// if err := json.Unmarshal(bytes, &user); err != nil {
	// 	fmt.Println(err.Error())
	// }

	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("Invalid json body")
		c.JSON(int(restError.Status), restError)
		return
	}
	result, errResult := services.CreateUser(user)
	if errResult != nil {
		c.JSON(int(errResult.Status), errResult)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("id should be a number !")
		c.JSON(int(err.Status), err)
		return
	}

	result, errResult := services.GetUser(userId)
	if errResult != nil {
		c.JSON(int(errResult.Status), errResult)
		return
	}
	c.JSON(http.StatusAccepted, result)
}

func PutUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("id should be a number !")
		c.JSON(int(err.Status), err)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("Invalid json body")
		c.JSON(int(restError.Status), restError)
		return
	}

	user.Id = userId
	result, err := services.UpdateUser(user)
	if err != nil {
		c.JSON(int(err.Status), err)
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("User id must be number")
		c.JSON(int(err.Status), err)
		return
	}
	var user users.User
	user.Id = userId
	if err := services.DeleteUser(user); err != nil {
		c.JSON(int(err.Status), err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func GetUsers(c *gin.Context) {
	result, errResult := services.GetUserCollection(c)
	if errResult != nil {
		c.JSON(int(errResult.Status), errResult)
		return
	}
	c.JSON(http.StatusAccepted, result)
}

func Search(c *gin.Context) {
	query := c.Query("status")
	if query == "" {
		c.JSON(200, "Not found any users")
		return
	}
	status, err := strconv.ParseBool(query)
	if err != nil {
		c.JSON(400, err)
		return
	}
	result, errResult := services.GetUserByStatus(status)
	if errResult != nil {
		c.JSON(int(errResult.Status), errResult)
		return
	}
	c.JSON(http.StatusAccepted, result)
}
