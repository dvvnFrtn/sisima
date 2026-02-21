package dto

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

type ValidationError struct {
	Errors map[string][]string `json:"errors"`
}

func (e *ValidationError) Error() string {
	return "validation failed"
}

func Validate(s interface{}) *ValidationError {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	err := validate.Struct(s)
	if err == nil {
		return nil
	}

	t := reflect.TypeOf(s).Elem()
	errors := make(map[string][]string)

	for _, err := range err.(validator.ValidationErrors) {
		fieldName := err.Field()

		jsonTag := ""
		if field, ok := t.FieldByName(fieldName); ok {
			jsonTag = field.Tag.Get("json")
			if jsonTag != "" && jsonTag != "-" {
				fieldName = jsonTag
			}
		}

		tag := err.Tag()
		var message string
		switch tag {
		case "required":
			message = "is required"
		case "min":
			message = "minimum is " + err.Param() + " characters"
		case "max":
			message = "maximum is " + err.Param() + " characters"
		case "len":
			message = "must be exactly " + err.Param() + " characters"
		case "numeric":
			message = "must be a number"
		case "alphaunicode":
			message = "must contain only letters"
		case "oneof":
			message = "must be one of: " + err.Param()
		case "gt":
			message = "must be greater than: " + err.Param()
		default:
			message = "is invalid"
		}

		errors[fieldName] = append(errors[fieldName], message)
	}

	return &ValidationError{Errors: errors}
}
