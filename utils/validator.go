package utils

import (
	"fmt"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
)

var once sync.Once

func GetValidator() *validator.Validate {
	var validate *validator.Validate
	once.Do(func() {
		validate = validator.New()
	})
	return validate
}

func FormatValidationError(err error) string {
	var errorMessages []string
	for _, e := range err.(validator.ValidationErrors) {
		// filter message
		errorMessages = append(errorMessages, fmt.Sprintf("%s is required", e.Field()))
	}
	return strings.Join(errorMessages, ", ")
}
