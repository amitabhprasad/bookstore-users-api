package mysql_utils

import (
	"strings"

	"github.com/amitabhprasad/bookstore-util-go/logger"
	"github.com/amitabhprasad/bookstore-util-go/rest_errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errNoRow = "no rows in result set"
)

func ParseError(err error) *rest_errors.RestErr {
	logger.Error("Error in database transaction", err)
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errNoRow) {
			return rest_errors.NewNotFoundError("no matching record found with given input ")
		}
		return rest_errors.NewInternalServerError("error parsing sql response", sqlErr)
	}
	switch sqlErr.Number {
	case 1062:
		return rest_errors.NewbadRequestError("invalid data: ")
	}
	return rest_errors.NewInternalServerError("error processing request ", sqlErr)
}
