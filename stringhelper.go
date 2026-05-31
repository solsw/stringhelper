package stringhelper

import (
	"strings"
)

// Insert inserts the values 'ss' into 's' at rune index 'i', returning the modified string.
// Insert panics if 'i' is out of range.
func Insert(s string, i int, ss ...string) string {
	rr := []rune(s)
	return string(rr[:i]) + strings.Join(ss, "") + string(rr[i:])
}

// SkipAny returns a copy of 's' with all Unicode code points contained in 'chars' removed.
func SkipAny(s, chars string) string {
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune(chars, r) {
			return -1
		}
		return r
	}, s)
}

// JoinSkip is like [strings.Join], but skips the elements for which 'skip' (if provided) returns true.
func JoinSkip(elems []string, sep string, skip func(string) bool) string {
	if skip == nil {
		return strings.Join(elems, sep)
	}
	var b strings.Builder
	first := true
	for _, elem := range elems {
		if skip(elem) {
			continue
		}
		if !first {
			b.WriteString(sep)
		}
		b.WriteString(elem)
		first = false
	}
	return b.String()
}

// ReplaceNewLines replaces end-of-line markers (see [bufio.ScanLines]) within 's' with 'repl'.
func ReplaceNewLines(s, repl string) string {
	// "\r\n" must be matched before "\n" so existing CRLF markers are
	// consumed whole and not corrupted into "\r" + repl.
	return strings.NewReplacer("\r\n", repl, "\n", repl).Replace(s)
}

// StringToStrings slices 's' into all substrings separated by end-of-line markers (see [bufio.ScanLines]).
// If 's' is empty, a slice with a single element - an empty string - is returned.
func StringToStrings(s string) []string {
	return strings.Split(ReplaceNewLines(s, "\n"), "\n")
}
