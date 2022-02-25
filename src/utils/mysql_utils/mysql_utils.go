package mysql_utils

import (
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/laithrafid/bookstore_oauth-api/src/utils/errors_utils"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) errors_utils.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return errors_utils.NewNotFoundError("no record matching given id")
		}
		return errors_utils.NewInternalServerError("error parsing database response", err)
	}

	switch sqlErr.Number {
	case 1062:
		return errors_utils.NewBadRequestError("invalid data")
	}
	return errors_utils.NewInternalServerError("error processing request", errors.New("database error"))
}
