package user

import (
	"fmt"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/date_utils"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/domain/users"
	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/errors"
)

func CreateUser(u users.User) (*users.User, *errors.RestErr) {
	fmt.Println("inside create user service ", u)
	if err := u.Validate(); err != nil {
		return nil, err
	}
	u.DateCreated = date_utils.GetNowDBFormat()
	u.Status = users.StatusActive
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

func UpdateUser(isPartial bool, u users.User) (*users.User, *errors.RestErr) {
	currentUser, err := GetUser(u.Id)
	if err != nil {
		return nil, err
	}
	// if err := u.Validate(); err != nil {
	// 	return nil, err
	// }
	if isPartial {
		currentUser.UpdateFields(&u)
	} else {
		currentUser.FirstName = u.FirstName
		currentUser.LastName = u.LastName
		if u.Email != "" {
			currentUser.Email = u.Email
		}
	}
	err = currentUser.Update()
	if err != nil {
		fmt.Println("ERROR ", err)
		return nil, err
	}
	return currentUser, nil
}

func DeleteUser(id int64) *errors.RestErr {
	user := &users.User{Id: id}
	return user.Delete()
}

func Search(status string) ([]users.User, *errors.RestErr) {
	user := &users.User{}
	return user.FindByStatus(status)
}
