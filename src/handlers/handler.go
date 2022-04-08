package handlers

import (
	"github.com/go-playground/validator/v10"
)

type _Return struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}

type _ReturnErr struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

func Success(message string, body interface{}) _Return {
	return _Return{
		Success: true,
		Message: message,
		Body:    body,
	}
}

func Error(err error) _ReturnErr {
	return _ReturnErr{
		Success: false,
		Error:   err.Error(),
	}
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return err
	}
	return nil
}
