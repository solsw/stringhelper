package stringhelper

import (
	"strings"
)

// IsEmptyOrWhite reports whether 's' is empty or contains only white spaces, as defined by Unicode.
func IsEmptyOrWhite(s string) bool {
	if s == "" {
		// with this 'if' the function is ~10 times faster for empty string
		return true
	}
	return strings.TrimSpace(s) == ""
}

// IsUpper reports whether 's' is upper case.
// A string containing no cased runes (e.g. "123") is considered both upper and lower case.
func IsUpper(s string) bool {
	return s == strings.ToUpper(s)
}

// IsLower reports whether 's' is lower case.
// A string containing no cased runes (e.g. "123") is considered both upper and lower case.
func IsLower(s string) bool {
	return s == strings.ToLower(s)
}
