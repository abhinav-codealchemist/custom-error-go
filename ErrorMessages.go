package error

var ErrorCodeToUserMessages = map[ErrorCode]string{
	REQUEST_NOT_VALID: "Invalid Request. Please try again.",
	SQL_INSERT_ERROR:  "Something went wrong. Please try again.",
	SQL_UPDATE_ERROR:  "Something went wrong. Please try again.",
	SQL_FETCH_ERROR:   "Something went wrong. Please try again.",
}
