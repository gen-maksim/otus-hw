package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type ParsedValidator struct {
	CondType string
	CondVal  string
}

type ValidationError struct {
	Field string
	Err   error
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Err.Error())
}

func (e ValidationError) Is(err error) bool {
	return errors.Is(err, e.Err)
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Error() string {
	result := strings.Builder{}
	result.WriteString("Validation Errors:\n")
	for _, err := range v {
		result.WriteString(err.Error())
	}

	return result.String()
}

func (v ValidationErrors) Unwrap() []error {
	if len(v) == 0 {
		return nil
	}
	result := []error{}
	for _, err := range v {
		result = append(result, err)
	}

	return result
}

func Validate(v interface{}) error {
	rt := reflect.TypeOf(v)
	if rt.Kind() != reflect.Struct {
		return nil
	}
	var errBag ValidationErrors

	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		alias, ok := field.Tag.Lookup("validate")
		if !ok || alias == "" {
			continue
		}

		err := validateField(field, reflect.ValueOf(v).Field(i), alias)
		if err != nil {
			var errT ValidationErrors
			if !errors.As(err, &errT) {
				return err
			}

			errBag = append(errBag, errT...)
		}
	}

	if len(errBag) > 0 {
		return errBag
	}

	return nil
}

func validateField(key reflect.StructField, v reflect.Value, validator string) error {
	if v.Kind() == reflect.Slice {
		var errBag ValidationErrors

		for i := 0; i < v.Len(); i++ {
			err := validateField(key, v.Index(i), validator)
			var errT ValidationErrors
			if !errors.As(err, &errT) {
				return err
			}

			errBag = append(errBag, errT...)
		}

		return errBag
	}

	validatorAr := []ParsedValidator{}
	for _, oneRawVal := range strings.Split(validator, "|") {
		oneVal := strings.Split(oneRawVal, ":")
		if len(oneVal) != 2 {
			continue
		}
		parsed := ParsedValidator{CondType: oneVal[0], CondVal: oneVal[1]}

		validatorAr = append(validatorAr, parsed)
	}

	if v.Kind() == reflect.String {
		return validateString(key, v, validatorAr)
	}

	if v.Kind() == reflect.Int {
		return validateInt(key, v, validatorAr)
	}

	return nil
}
