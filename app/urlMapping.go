package app

import (
	ping "bookStoreUser/controllers/ping"
	users "bookStoreUser/controllers/users"
)

func mapUrl() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:id", users.GetUser)
	router.GET("/users", users.GetUsers)
	router.PUT("/users/:id", users.PutUser)
	router.DELETE("/users/:id", users.DeleteUser)
}
