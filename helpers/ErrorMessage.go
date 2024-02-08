package helpers

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ErrorMessage struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func ValidationErrorMessage(err error, status int) *ErrorMessage {
	var ve validator.ValidationErrors
	if !errors.As(err, &ve) {
		return &ErrorMessage{
			Status:  500,
			Message: "Unknown Error",
		}
	}
	return &ErrorMessage{
		Status:  status,
		Message: strings.ToLower(ve[0].Field()) + " " + ve[0].Tag(),
	}
}
