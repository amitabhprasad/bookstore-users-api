package app

import (
	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/controllers/ping"
	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/controllers/users"
)

func mapUrls() {
	// function is not executed but passed as object
	router.GET("/ping", ping.Ping)
	router.GET("/users/search/:user_name", users.SearchUser)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)

}
