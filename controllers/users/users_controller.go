package users

import (
	"net/http"
	"strconv"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/domain/users"
	service_user "github.com/amitabhprasad/bookstore-app/bookstore-users-api/services/users"
	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewbadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, err := service_user.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		restErr := errors.NewbadRequestError("Invalid userid, userid should be number")
		c.JSON(restErr.Status, restErr)
		return
	}

	user, getErr := service_user.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me SearchUser !")
}
