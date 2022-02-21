package users

import (
	"net/http"
	"strconv"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/domain/users"
	service_user "github.com/amitabhprasad/bookstore-app/bookstore-users-api/services/users"
	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func getUserID(c *gin.Context) (int64, *errors.RestErr) {
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return 0, errors.NewbadRequestError("Invalid userid, userid should be number")
	}
	return userId, nil
}
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
	userId, idError := getUserID(c)
	if idError != nil {
		c.JSON(idError.Status, idError)
		return
	}

	user, getErr := service_user.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	userId, idError := getUserID(c)
	if idError != nil {
		c.JSON(idError.Status, idError)
		return
	}

	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewbadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := service_user.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusFound, result)
}

func DeleteUser(c *gin.Context) {
	userId, idError := getUserID(c)
	if idError != nil {
		c.JSON(idError.Status, idError)
		return
	}
	deleteErr := service_user.DeleteUser(userId)
	if deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := service_user.Search(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users)
}
