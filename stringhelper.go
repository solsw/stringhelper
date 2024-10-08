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
	var b strings.Builder
	for _, elem := range elems {
		if skip != nil && skip(elem) {
			continue
		}
		if b.Len() > 0 {
			b.WriteString(sep)
		}
		b.WriteString(elem)
	}
	return b.String()
}

// ReplaceNewLines replaces end-of-line markers (see [bufio.ScanLines]) within 's' with 'new'.
func ReplaceNewLines(s, new string) string {
	var r *strings.Replacer
	switch new {
	case "\n":
		r = strings.NewReplacer("\r\n", new)
	case "\r\n":
		r = strings.NewReplacer("\n", new)
	default:
		r = strings.NewReplacer("\r\n", new, "\n", new)
	}
	return r.Replace(s)
}

// StringToStrings slices 's' into all substrings separated by end-of-line markers (see [bufio.ScanLines]).
// If 's' is empty, a slice with a single element - an empty string - is returned.
func StringToStrings(s string) []string {
	return strings.Split(ReplaceNewLines(s, "\n"), "\n")
}
