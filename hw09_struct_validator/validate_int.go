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

func validateInt(key reflect.StructField, s reflect.Value, validatorAr []ParsedValidator) error {
	var errBag ValidationErrors
	for _, oneVal := range validatorAr {
		var err error

		switch oneVal.CondType {
		case "in":
			err = validateIn(strconv.FormatInt(s.Int(), 10), oneVal.CondVal)
		case "min":
			minV, sysErr := strconv.Atoi(oneVal.CondVal)
			if sysErr != nil {
				return sysErr
			}
			if int64(minV) >= s.Int() {
				err = fmt.Errorf("%w %v", ErrorMin, oneVal.CondVal)
			}
		case "max":
			minV, sysErr := strconv.Atoi(oneVal.CondVal)
			if sysErr != nil {
				return sysErr
			}
			if int64(minV) <= s.Int() {
				err = fmt.Errorf("%w %v", ErrorMax, oneVal.CondVal)
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
