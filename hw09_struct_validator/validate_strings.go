package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

var (
	ErrorIn    = errors.New("field should be in")
	ErrorLen   = errors.New("length of field should be equal to")
	ErrorRegex = errors.New("field should match pattern")
)

func validateString(key reflect.StructField, s reflect.Value, validatorAr []ParsedValidator) ValidationErrors {
	var errBag ValidationErrors
	for _, oneVal := range validatorAr {
		var err error

		switch oneVal.CondType {
		case "in":
			err = validateIn(s.String(), oneVal.CondVal)
		case "len":
			err = validateLen(s.String(), oneVal.CondVal)
		case "regexp":
			err = validateRegex(s.String(), oneVal.CondVal)
		}

		if err != nil {
			errBag = append(errBag, ValidationError{Field: key.Name, Err: err})
		}
	}

	if len(errBag) > 0 {
		return errBag
	}

	return nil
}

func validateIn(n string, rawValidators string) error {
	haystack := strings.Split(rawValidators, ",")
	if slices.Contains(haystack, n) {
		return nil
	}

	return fmt.Errorf("%w %s", ErrorIn, rawValidators)
}

func validateLen(n string, length string) error {
	l, err := strconv.Atoi(length)
	if err != nil {
		return err
	}

	if len(n) != l {
		return fmt.Errorf("%w %s", ErrorLen, length)
	}

	return nil
}

func validateRegex(n string, regex string) error {
	match, err := regexp.Match(regex, []byte(n))
	if err != nil {
		return err
	}
	if !match {
		return fmt.Errorf("%w %s", ErrorRegex, regex)
	}

	return nil
}
