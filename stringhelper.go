package stringhelper

import (
	"strings"
)

// JoinSkip is like strings.Join, but skips the elements for which 'skip' (if any) returns true.
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
