package hw09structvalidator

import (
	"errors"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func validateString(key reflect.StructField, s reflect.Value, rawValidators string) ValidationErrors {
	validatorAr := strings.Split(rawValidators, "|")
	var errBag ValidationErrors
	for _, oneRawVal := range validatorAr {
		oneVal := strings.Split(oneRawVal, ":")
		var err error

		switch oneVal[0] {
		case "in":
			err = validateIn(s.String(), oneVal[1])
		case "len":
			err = validateLen(s.String(), oneVal[1])
		case "regexp":
			err = validateRegex(s.String(), oneVal[1])
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

	return errors.New("value should be in " + rawValidators)
}

func validateLen(n string, length string) error {
	l, err := strconv.Atoi(length)
	if err != nil {
		return err
	}

	if len(n) > l {
		return errors.New("length should be less than or equal to " + length)
	}

	return nil
}

func validateRegex(n string, regex string) error {
	match, err := regexp.Match(regex, []byte(n))
	if err != nil {
		return err
	}
	if !match {
		return errors.New("value should match pattern " + regex)
	}

	return nil
}
