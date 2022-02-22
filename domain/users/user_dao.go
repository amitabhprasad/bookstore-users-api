package users

// dao access layer to database
import (
	"fmt"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/logger"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/mysql_utils"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/datasources/mysql/users_db"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/errors"
)

var (
	userDB     = make(map[int64]*User)
	errorNoRow = "no rows in result set"
)

const (
	queryInsertUser       = "INSERT into users (first_name,last_name,email,date_created,password,status) VALUES (?,?,?,?,?,?);"
	queryGetUserById      = "SELECT id, first_name, last_name, email, date_created, status from users WHERE id = ?;"
	queryUpdateUser       = "UPDATE users set first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE from users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status = ?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUserById)
	if err != nil {
		logger.Error("error during prepare statement for get user", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		return mysql_utils.ParseError(getErr)
	}
	return nil
}
func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error during prepare statement for save user", err)
		return errors.NewInternalServerError("unable to save user info in database")
	}
	defer stmt.Close()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Password, user.Status)
	if saveErr != nil {
		return mysql_utils.ParseError(saveErr)
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return mysql_utils.ParseError(err)
	}
	user.Id = userId
	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error during prepare statement for update user", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_, updateErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if updateErr != nil {
		return mysql_utils.ParseError(updateErr)
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error during prepare statement for delete user", err)
		return errors.NewInternalServerError("database error, unable to delete user")
	}
	defer stmt.Close()
	_, delError := stmt.Exec(user.Id)
	if delError != nil {
		return mysql_utils.ParseError(delError)
	}
	return nil
}

func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error during prepare statement for find users", err)
		return nil, errors.NewInternalServerError("database error during find users")
	}
	defer stmt.Close()
	rows, findErr := stmt.Query(status)
	if findErr != nil {
		return nil, mysql_utils.ParseError(findErr)
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysql_utils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		// shouldn't log this message since this is due to user error and not because of application function
		return nil, errors.NewNotFoundError(fmt.Sprintf("No user matching given status %s ", status))
	}
	return results, nil
}
