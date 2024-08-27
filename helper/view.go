package helper

import (
	"strings"
	"unicode"
)

func ToLower(str string) string {
	return strings.ToLower(str)
}

func CapitalizeFirstLetter(str string) string {
	if len(str) == 0 {
		return ""
	}

	firstLetter := string(unicode.ToUpper(rune(str[0])))
	rest := strings.ToLower(str[1:])

	return firstLetter + rest
}

func CapitalizeAfterPeriod(s string) string {
	if len(s) == 0 {
		return ""
	}

	// Convert the entire string to lowercase first
	result := strings.ToLower(s)

	// Convert the first character to uppercase
	runes := []rune(result)
	runes[0] = unicode.ToUpper(runes[0])

	// Iterate through the string and capitalize any letter that comes after a period
	for i := 1; i < len(runes); i++ {
		if runes[i-1] == '.' && i+1 < len(runes) && unicode.IsLetter(runes[i]) {
			runes[i] = unicode.ToUpper(runes[i])
		}
	}

	return string(runes)
}
