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
		if prev > 0 { //not first
			if ok == nil { //current is number
				if prevok == nil { //prev is number too
					return "", ErrInvalidString
				}

				res += strings.Repeat(string(prev), repeat)
			} else { //current is letter
				if prevok != nil { //prev is letter
					res += string(prev)
				}
				if index == utf8.RuneCountInString(s)-1 { //curr is last
					res += string(value)
				}
			}
		} else if ok == nil {
			return "", ErrInvalidString
		}

		prev = value
		index++
	}

	return res, nil
}
