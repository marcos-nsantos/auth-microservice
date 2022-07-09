package helper

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/non-standard/validators"
	"log"
)

var validate *validator.Validate

func RegisterValidators() {
	validate = validator.New()
	if err := validate.RegisterValidation("notblank", validators.NotBlank); err != nil {
		log.Panicf("failed to register notblank validator: %v", err)
	}
}

func getErrorValidationMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", fe.Field())
	case "notblank":
		return fmt.Sprintf("%s cannot be blank", fe.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s caracters", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s caracters", fe.Field(), fe.Param())
	case "email":
		return fmt.Sprintf("%s is not valid", fe.Field())
	}
	return fmt.Sprintf("%s is invalid", fe.Field())
}

func CheckForValidationErrMessages(s interface{}) error {
	err := validate.Struct(s)
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, fe := range ve {
			return fmt.Errorf("%s", getErrorValidationMessage(fe))
		}
	}
	return err
}
