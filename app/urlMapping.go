package app

import (
	ping "bookStoreUser/controllers/ping"
	users "bookStoreUser/controllers/users"
)

func mapUrl() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:id", users.GetUser)
}
