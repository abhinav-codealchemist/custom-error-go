package error

import "fmt"

type ErrorCode string

func (e ErrorCode) ToString() string {
	return fmt.Sprintf("%s", e)
}

const (
	REQUEST_NOT_VALID ErrorCode = "REQUEST_NOT_VALID"
	SQL_INSERT_ERROR  ErrorCode = "SQL_INSERT_ERROR"
	SQL_UPDATE_ERROR  ErrorCode = "SQL_UPDATE_ERROR"
	SQL_FETCH_ERROR   ErrorCode = "SQL_FETCH_ERROR"
)
