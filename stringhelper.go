package stringhelper

import (
	"strings"
)

// JoinSkip is like [strings.Join], but skips the elements for which 'skip' (if any) returns true.
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
// If 's' is empty, a slice with only element - empty string is returned.
func StringToStrings(s string) []string {
	return strings.Split(ReplaceNewLines(s, "\n"), "\n")
}
