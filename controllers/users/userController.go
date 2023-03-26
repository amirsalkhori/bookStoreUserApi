package controllers

import (
	"bookStoreUser/domain/users"
	"bookStoreUser/errors"
	"bookStoreUser/services"
	"net/http"

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

func GetUser() {

}

func FindUser() {

}
