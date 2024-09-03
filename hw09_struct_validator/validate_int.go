package hw09structvalidator

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

func validateInt(key reflect.StructField, s reflect.Value, rawValidators string) ValidationErrors {
	validatorAr := strings.Split(rawValidators, "|")
	var errBag ValidationErrors
	for _, oneRawVal := range validatorAr {
		oneVal := strings.Split(oneRawVal, ":")
		var err error

		switch oneVal[0] {
		case "in":
			err = validateIn(strconv.FormatInt(s.Int(), 10), oneVal[1])
		case "min":
			minV, _ := strconv.Atoi(oneVal[1])
			if int64(minV) >= s.Int() {
				err = errors.New("value should be more or equal " + oneVal[1])
			}
		case "max":
			minV, _ := strconv.Atoi(oneVal[1])
			if int64(minV) <= s.Int() {
				err = errors.New("value should be less or equal " + oneVal[1])
			}
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
