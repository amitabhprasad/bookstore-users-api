package user

import (
	"fmt"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/domain/users"
	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/errors"
)

func CreateUser(u users.User) (*users.User, *errors.RestErr) {
	fmt.Println("inside create user service")
	if err := u.Validate(); err != nil {
		return nil, err
	}
	if err := u.Save(); err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}
