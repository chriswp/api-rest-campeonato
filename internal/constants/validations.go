package constants

import "fmt"

const RequiredFieldErrorMessage = "'%s' is required"
const InvalidFieldErrorMessage = "invalid '%s' format"
const IsAlreadyExistsMessage = "'%s' is already exists"

func RequiredField(field string) string {
	return fmt.Sprintf(RequiredFieldErrorMessage, field)
}

func InvalidFieldError(field string) string {
	return fmt.Sprintf(InvalidFieldErrorMessage, field)
}

func IsAlreadyExists(field string) string {
	return fmt.Sprintf(IsAlreadyExistsMessage, field)
}
