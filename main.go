package main

import (
	"fmt"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/app"
)

func main() {
	fmt.Println("Executing module ... bookstore user API")
	//loadEnvVariables()
	app.StartApplication()
	// r := gin.Default()
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run("0.0.0.0:8081")
}
