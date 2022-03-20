package services

import (
	"fmt"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/crypto_utils"
	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/date_utils"
	"github.com/amitabhprasad/bookstore-util-go/rest_errors"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/domain/users"
)

var (
	UsersService userServiceInterface = &usersService{}
)

type usersService struct {
}

type userServiceInterface interface {
	GetUser(int64) (*users.User, *rest_errors.RestErr)
	CreateUser(users.User) (*users.User, *rest_errors.RestErr)
	UpdateUser(bool, users.User) (*users.User, *rest_errors.RestErr)
	DeleteUser(int64) *rest_errors.RestErr
	SearchUser(string) (users.Users, *rest_errors.RestErr)
	LoginUser(users.UserLoginRequest) (*users.User, *rest_errors.RestErr)
}

func (s *usersService) CreateUser(u users.User) (*users.User, *rest_errors.RestErr) {
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

func (s *usersService) GetUser(userId int64) (*users.User, *rest_errors.RestErr) {
	user := &users.User{Id: userId}
	if err := user.Get(); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *usersService) UpdateUser(isPartial bool, u users.User) (*users.User, *rest_errors.RestErr) {
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

func (s *usersService) DeleteUser(id int64) *rest_errors.RestErr {
	user := &users.User{Id: id}
	return user.Delete()
}

func (s *usersService) SearchUser(status string) (users.Users, *rest_errors.RestErr) {
	user := &users.User{}
	return user.FindByStatus(status)
}

func (s *usersService) LoginUser(request users.UserLoginRequest) (*users.User, *rest_errors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: crypto_utils.GetMD5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, rest_errors.NewNotFoundError(
			fmt.Sprintf("Invalid user credentials for user %s ", request.Email))
	}
	return dao, nil
}
