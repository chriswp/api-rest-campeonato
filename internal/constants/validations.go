package constants

import "fmt"

const RequiredFieldError = "o campo '%s' é obrigatório"

func FormatRequiredField(field string) string {
	return fmt.Sprintf(RequiredFieldError, field)
}
