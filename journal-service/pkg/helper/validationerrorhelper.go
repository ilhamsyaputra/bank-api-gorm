package helper

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func RequestValidation(err error) (code int, message string) {
	// https://dasarpemrogramangolang.novalagung.com/C-http-error-handling.html
	code = fiber.ErrBadRequest.Code
	var fieldRequired []string

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				field := fmt.Sprintf("%v,", toSnakeCase(err.Field()))
				fieldRequired = append(fieldRequired, field)
			}
		}

		if len(fieldRequired) > 0 {
			message = fmt.Sprintf("REQUIRED: field %v harus diisi. ", fieldRequired)
		}
	}
	return
}

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
