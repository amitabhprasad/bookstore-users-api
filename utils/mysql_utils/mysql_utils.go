package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/logger"
	"github.com/amitabhprasad/bookstore-app/bookstore-users-api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errNoRow = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	logger.Error("Error in database transaction", err)
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errNoRow) {
			return errors.NewNotFoundError(fmt.Sprintf("no matching record found with given input "))
		}
		return errors.NewInternalServerError("error parsing sql response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewbadRequestError(fmt.Sprintf("invalid data: "))
	}
	return errors.NewInternalServerError(fmt.Sprintf("error processing request "))
}
