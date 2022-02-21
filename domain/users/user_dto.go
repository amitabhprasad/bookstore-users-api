package users

import (
	"strings"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/errors"
)

const (
	StatusActive = "active"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Password = strings.TrimSpace(user.Password)
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	// email address is automatically cleaned i.e. removed of white spaces and lowercased and validated
	if user.Email == "" {
		return errors.NewbadRequestError("invalid email address")
	}
	if user.Password == "" {
		return errors.NewbadRequestError("invalid password ")
	}
	return nil
}

func (user *User) UpdateFields(u *User) {
	if u.Email != "" {
		user.Email = u.Email
	}
	if u.FirstName != "" {
		user.FirstName = u.FirstName
	}
	if u.LastName != "" {
		user.LastName = u.LastName
	}
}
