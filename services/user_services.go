package services

import (
	"fmt"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/crypto_utils"
	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/date_utils"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/domain/users"
	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/errors"
)

var (
	UsersService userServiceInterface = &usersService{}
)

type usersService struct {
}

type userServiceInterface interface {
	GetUser(int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
}

func (s *usersService) CreateUser(u users.User) (*users.User, *errors.RestErr) {
	fmt.Println("inside create user service ", u)
	if err := u.Validate(); err != nil {
		return nil, err
	}
	u.DateCreated = date_utils.GetNowDBFormat()
	u.Password = crypto_utils.GetMD5(u.Password)
	u.Status = users.StatusActive
	if err := u.Save(); err != nil {
		return nil, err
	}
	return &u, nil
}

func (s *usersService) GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *usersService) UpdateUser(isPartial bool, u users.User) (*users.User, *errors.RestErr) {
	currentUser, err := s.GetUser(u.Id)
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

func (s *usersService) DeleteUser(id int64) *errors.RestErr {
	user := &users.User{Id: id}
	return user.Delete()
}

func (s *usersService) SearchUser(status string) (users.Users, *errors.RestErr) {
	user := &users.User{}
	return user.FindByStatus(status)
}