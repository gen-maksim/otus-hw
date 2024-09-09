package hw09structvalidator

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

var (
	ErrorMin = errors.New("field should be more or equal")
	ErrorMax = errors.New("field should be less or equal")
)

func validateInt(key reflect.StructField, s reflect.Value, validatorAr [][]string) error {
	var errBag ValidationErrors
	for _, oneVal := range validatorAr {
		var err error

		switch oneVal[0] {
		case "in":
			err = validateIn(strconv.FormatInt(s.Int(), 10), oneVal[1])
		case "min":
			minV, sysErr := strconv.Atoi(oneVal[1])
			if sysErr != nil {
				return sysErr
			}
			if int64(minV) >= s.Int() {
				err = fmt.Errorf("%w %v", ErrorMin, oneVal[1])
			}
		case "max":
			minV, sysErr := strconv.Atoi(oneVal[1])
			if sysErr != nil {
				return sysErr
			}
			if int64(minV) <= s.Int() {
				err = fmt.Errorf("%w %v", ErrorMax, oneVal[1])
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
