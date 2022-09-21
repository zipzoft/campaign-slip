package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var _ error = (*ApiError)(nil)

type ApiError struct {
	Field   string
	Message string
}

func New(field string, message string) error {
	return &ApiError{field, message}
}

func (a *ApiError) Error() string {
	return a.Message
}

func Validate(model interface{}) []ApiError {
	validate := validator.New()

	err := validate.Struct(model)

	if err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ApiError, len(ve))
			for i, fe := range ve {
				out[i] = ApiError{fe.Field(), msgForTag(fe.Tag(), fe.Field())}
			}
			return out
		}
	}

	return nil
}

func msgForTag(tag string, field string) string {
	switch tag {
	case "required":
		return field + " field is required"
	case "gte":
		return "This data must be array"
	}
	return ""
}
