package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errNoRow = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	fmt.Println("Processing error ", err)
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errNoRow) {
			return errors.NewNotFoundError(fmt.Sprintf("no matching record found with given input %s", err.Error()))
		}
		return errors.NewInternalServerError("error parsing sql response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewbadRequestError(fmt.Sprintf("invalid data:  %s ", sqlErr.Message))
	}
	return errors.NewInternalServerError(fmt.Sprintf("error processing request %s", err))
}
