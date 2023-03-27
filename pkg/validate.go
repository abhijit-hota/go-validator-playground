package pkg

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type InvalidResult struct {
	Fields []string `json:"fields"`
	Tags   []string `json:"tags"`
}

func ValidateStruct(foo any) (res *InvalidResult, err error) {
	v := validator.New()

	defer func() {
		if r := recover(); r != nil {
			err = errors.New("panic occurred: please check your struct tags")
		}
	}()

	if err := v.Struct(foo); err != nil {
		if invalidErr, ok := err.(*validator.InvalidValidationError); ok {
			return nil, invalidErr
		}

		res = ParseValidationError(err)
	}

	return
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
