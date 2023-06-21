package controllers

import (
	"bookStoreUser/domain/users"
	"bookStoreUser/errors"
	"bookStoreUser/services"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/amirsalkhori/bookstroe_oauth_go/oauth"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restError := errors.NewBadRequestError("Invalid json body")
		c.JSON(int(restError.Status), restError)
		return
	}
	result, errResult := services.UserService.CreateUser(user)
	if errResult != nil {
		c.JSON(int(errResult.Status), errResult)
		return
	}

	header := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusCreated, result.Marshall(header))
}

func GetUser(c *gin.Context) {

	baseURL := "http://localhost:8080"
	accessTokenID := "ali1"

	client := http.Client{
		Timeout: 100 * time.Millisecond,
	}

	url := fmt.Sprintf("%s/oauth/access_token/%s", baseURL, accessTokenID)
	response, err := client.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("Response is:", string(body))

	return

	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(int(err.Status), err)
		return
	}
	userId, userErr := strconv.ParseInt(c.Param("id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("id should be a number !")
		c.JSON(int(err.Status), err)
		return
	}

	result, errResult := services.UserService.GetUser(userId)
	if errResult != nil {
		c.JSON(int(errResult.Status), errResult)
		return
	}

	if oauth.GetCallerId(c.Request) == int(result.Id) {
		c.JSON(http.StatusAccepted, result.Marshall(false))
		return
	}
	// header := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusAccepted, result.Marshall(oauth.IsPublic(c.Request)))
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
	result, err := services.UserService.UpdateUser(user)
	if err != nil {
		c.JSON(int(err.Status), err)
		return
	}
	header := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusOK, result.Marshall(header))
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
	if err := services.UserService.DeleteUser(user); err != nil {
		c.JSON(int(err.Status), err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func GetUsers(c *gin.Context) {
	users, errResult := services.UserService.GetUserCollection(c)
	if errResult != nil {
		c.JSON(int(errResult.Status), errResult)
		return
	}
	// c.JSON(http.StatusAccepted, result)
	header := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusOK, users.Marshall(header))
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
	users, errResult := services.UserService.GetUserByStatus(status)
	if errResult != nil {
		c.JSON(int(errResult.Status), errResult)
		return
	}
	header := c.GetHeader("X-Public") == "true"
	c.JSON(http.StatusAccepted, users.Marshall(header))
}

func Login(c *gin.Context) {
	var request users.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restError := errors.NewBadRequestError("Invalid json body")
		c.JSON(int(restError.Status), restError)
		return
	}
	result, errResult := services.UserService.LoginUser(request)
	if errResult != nil {
		c.JSON(int(errResult.Status), errResult)
		return
	}

	c.JSON(http.StatusOK, result)
}
