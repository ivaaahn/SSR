package errors

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
)

//var (
//ErrNotFound      = NewAppError(nil, "Not found", 404)
//ErrNotAuthorized = NewAppError(nil, "Not authorized", 401)
//)

type AppError struct {
	Err        error  `json:"from-error,omitempty"`
	DevMessage string `json:"dev-message,omitempty"`
	Message    string `json:"message,omitempty"`
	Code       int    `json:"code,omitempty"`
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) ToHTTP() *echo.HTTPError {
	return echo.NewHTTPError(e.Code)
}

func (e *AppError) Marshal() []byte {
	marshal, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return marshal
}

func NewAppError(err error, devMsg, msg string, code int) *AppError {
	return &AppError{
		Err:        err,
		DevMessage: devMsg,
		Message:    msg,
		Code:       code,
	}
}
