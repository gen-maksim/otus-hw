package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var prev rune
	var res string
	var index int
	for _, value := range s {
		repeat, ok := strconv.Atoi(string(value))
		_, prevok := strconv.Atoi(string(prev))
		if prev > 0 { // not first
			if ok == nil && prevok == nil { // two numbers
				return "", ErrInvalidString
			}
			if ok == nil { // current is number
				res += strings.Repeat(string(prev), repeat)
			} else if prevok != nil { // prev is letter
				res += string(prev)
			}
		} else if ok == nil { // first is number
			return "", ErrInvalidString
		}

		if index == utf8.RuneCountInString(s)-1 && ok != nil { // curr is last
			res += string(value)
		}

		prev = value
		index++
	}

	return res, nil
}
