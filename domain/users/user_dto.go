package users

import (
	"strings"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	// email address is automatically cleaned i.e. removed of white spaces and lowercased and validated
	if user.Email == "" {
		return errors.NewbadRequestError("invalid email address")
	}
	return nil
}
