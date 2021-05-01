package error

import (
	"errors"
	"fmt"
	"strings"

	"github.com/Zomato/go/logger"
	errors2 "github.com/pkg/errors"
)

type CustomError struct {
	errorCode     ErrorCode
	errorMsg      string
	error         error
	exists        bool
	loggingParams map[string]interface{}
}

func NewCustomError(errorCode ErrorCode, error string) CustomError {
	c := CustomError{errorCode: errorCode, errorMsg: error, exists: true}
	e := errors.New(fmt.Sprintf("Code: %s | %s", c.errorCode, c.errorMsg))
	c.error = errors2.WithStack(e)
	c.loggingParams = make(map[string]interface{}, 0)
	return c
}

func (c CustomError) Exists() bool {
	return c.exists
}

func (c CustomError) Log() {
	logger.Errorln(c.ToString())
}

func (c CustomError) LoggingParams() map[string]interface{} {
	return c.loggingParams
}

func (c CustomError) ErrorCode() ErrorCode {
	return c.errorCode
}

func (c CustomError) ToError() error {
	return c.error
}

func (c CustomError) Error() string {
	return c.error.Error()
}

func (c CustomError) ErrorMessage() string {
	return c.errorMsg
}

func (c CustomError) ToString() string {
	logMsg := fmt.Sprintf("Code: %s, Msg: %s", c.errorCode, c.errorMsg)

	paramStrings := make([]string, 0)
	for key, val := range c.loggingParams {
		paramStrings = append(paramStrings, fmt.Sprintf("%s: {%+v}", strings.ToUpper(key), val))
	}
	return fmt.Sprintf("%s, Params: [%+v]", logMsg, strings.Join(paramStrings, " | "))
}

// value param should not be a pointer
func (c CustomError) WithParam(key string, val interface{}) CustomError {
	if c.loggingParams == nil {
		c.loggingParams = make(map[string]interface{}, 0)
	}
	c.loggingParams[key] = val
	return c
}

func (c CustomError) ErrorString() string {
	return c.errorMsg
}

func (c CustomError) UserMessage() string {
	if val, found := ErrorCodeToUserMessages[c.errorCode]; found {
		return val
	}
	return ""
}
