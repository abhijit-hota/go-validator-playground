package pkg

import (
	"github.com/go-playground/validator/v10"
)

type InvalidResult struct {
	Fields []string `json:"fields"`
	Tags   []string `json:"tags"`
}

func ValidateStruct(foo any) (*InvalidResult, error) {
	v := validator.New()

	if err := v.Struct(foo); err != nil {
		if invalidErr, ok := err.(*validator.InvalidValidationError); ok {
			return nil, invalidErr
		}

		return ParseValidationError(err), nil
	}

	return nil, nil
}

func ParseValidationError(err error) *InvalidResult {
	fields := make([]string, 0)
	tags := make([]string, 0)

	for _, validationErr := range err.(validator.ValidationErrors) {
		fields = append(fields, validationErr.Field())
		tags = append(tags, validationErr.Tag())
	}

	return &InvalidResult{
		Fields: fields,
		Tags:   tags,
	}
}
