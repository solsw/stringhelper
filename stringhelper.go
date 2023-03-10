package stringhelper

import (
	"bufio"
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

// StringToStrings slices 's' into all substrings separated by end-of-line markers (see [bufio.ScanLines]).
// If 's' is empty, empty slice is returned.
func StringToStrings(s string) []string {
	ss := []string{}
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		ss = append(ss, sc.Text())
	}
	return ss
}
