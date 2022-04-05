package app

import (
	"fmt"

	"github.com/amitabhprasad/bookstore-util-go/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
	port   = "8081"
)

func StartApplication() {
	mapUrls()
	logger.Info(fmt.Sprintf("Starting application...on port %s", port))
	router.Run(":" + port)
}
