package lib

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(validate *validator.Validate, v interface{}, msg string) (error, string) {
	if errs := validate.Struct(v); errs != nil {
		errSb := strings.Builder{}
		errSb.WriteString(msg)
		for _, err := range errs.(validator.ValidationErrors) {
			errSb.WriteString(fmt.Sprintf("%s - %s, ", err.Field(), err.Tag()))
		}

		return errs, errSb.String()
	}
	return nil, ""
}
