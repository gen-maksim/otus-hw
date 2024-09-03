package hw09structvalidator

import (
	"fmt"
	"reflect"
)

type ValidationError struct {
	Field string
	Err   error
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	result := "Validation Errors:\n"
	for _, err := range v {
		result += fmt.Sprintf("\t%s: %s\n", err.Field, err.Err)
	}
	return result
}

func Validate(v interface{}) error {
	rt := reflect.TypeOf(v)
	var errBag ValidationErrors

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if alias, ok := field.Tag.Lookup("validate"); ok {
			if alias != "" {
				err := validateField(field, reflect.ValueOf(v).Field(i), alias)
				if err != nil {
					errBag = append(errBag, err...)
				}
			}
		}
	}

	if len(errBag) > 0 {
		return errBag
	}

	return nil
}

func validateField(key reflect.StructField, v reflect.Value, validator string) ValidationErrors {
	if v.Kind() == reflect.String {
		return validateString(key, v, validator)
	}

	if v.Kind() == reflect.Int {
		return validateInt(key, v, validator)
	}

	return nil
}
