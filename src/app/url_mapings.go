package app

import (
	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/src/controllers/ping"
	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/src/controllers/users"
)

func mapUrls() {
	// function is not executed but passed as object
	router.GET("/ping", ping.Ping)
	router.GET("/users/search/", users.Search)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	router.POST("/users/login", users.LoginUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.PATCH("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)

}
