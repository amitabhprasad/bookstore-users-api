package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/domain/users"
	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/services"
	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/errors"

	"github.com/amitabhprasad/bookstore-oauth-go/oauth"
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
	result, err := services.UsersService.CreateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func GetUser(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		fmt.Println("Error ", err)
		//c.JSON(err.Status, err)
		//return
	}
	userId, idError := getUserID(c)
	if idError != nil {
		c.JSON(idError.Status, idError)
		return
	}

	user, getErr := services.UsersService.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}

	if oauth.GetCallerId(c.Request) == userId {
		c.JSON(http.StatusOK, user.Marshall(false))
		return
	}

	c.JSON(http.StatusOK, user.Marshall(oauth.IsPublic(c.Request)))
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

	result, updateErr := services.UsersService.UpdateUser(isPartial, user)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusFound, result.Marshall(c.GetHeader("X-Public") == "true"))
}

func DeleteUser(c *gin.Context) {
	userId, idError := getUserID(c)
	if idError != nil {
		c.JSON(idError.Status, idError)
		return
	}
	deleteErr := services.UsersService.DeleteUser(userId)
	if deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall(c.GetHeader("X-Public") == "true"))
}

func LoginUser(c *gin.Context) {
	var request users.UserLoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		restErr := errors.NewbadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, loginError := services.UsersService.LoginUser(request)
	if loginError != nil {
		c.JSON(loginError.Status, loginError)
		return
	}
	c.JSON(http.StatusOK, user.Marshall(c.GetHeader("X-Public") == "true"))
}
