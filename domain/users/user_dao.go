package users

// dao access layer to database
import (
	"fmt"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/errors"
)

var (
	userDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := userDB[user.Id]
	fmt.Println("called GET ", result)
	if result == nil {
		fmt.Println("inside if block....")
		return errors.NewNotFoundError(fmt.Sprintf("User with the id %d  not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}
func (user *User) Save() *errors.RestErr {
	current := userDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewUserExistsError(fmt.Sprintf("User with the email %s  already registered ", user.Email))
		}
		return errors.NewUserExistsError(fmt.Sprintf("User with the id %d  already exists ", user.Id))
	}
	userDB[user.Id] = user
	return nil
}
