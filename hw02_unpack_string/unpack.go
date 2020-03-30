package hw02_unpack_string //nolint:golint,stylecheck

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result strings.Builder
	var runes = []rune(str)
	for i, v := range runes {
		if unicode.IsDigit(runes[0]) {
			return "", ErrInvalidString
		}
		if !unicode.IsDigit(v) {
			result.WriteRune(v)
		} else {
			digit, _ := strconv.Atoi(string(v))
			if digit != 0 {
				result.WriteString(strings.Repeat(string(runes[i-1]), digit-1))
			} else {
				return "", ErrInvalidString
			}
		}
	}
	return result.String(), nil
}
