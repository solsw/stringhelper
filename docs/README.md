# stringhelper
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/stringhelper.svg)](https://pkg.go.dev/github.com/solsw/stringhelper)
[![GitHub](https://img.shields.io/badge/github--green?logo=github)](https://github.com/solsw/stringhelper)

Helpers for Go's [string](https://go.dev/ref/spec#String_types).

## Installation

```sh
go get github.com/solsw/stringhelper
```

```go
import "github.com/solsw/stringhelper"
```

## API

### Manipulation

#### `Insert(s string, i int, ss ...string) string`

Inserts the values `ss` into `s` at rune index `i`, returning the modified string.
Indexing is by rune, not by byte, so multi-byte characters are handled correctly.
Panics if `i` is out of range.

```go
stringhelper.Insert("acd", 1, "b")        // "abcd"
stringhelper.Insert("13", 1, "2")          // "123"
stringhelper.Insert("aé", 1, "X", "Y")     // "aXYé"
```

#### `SkipAny(s, chars string) string`

Returns a copy of `s` with all Unicode code points contained in `chars` removed.

```go
stringhelper.SkipAny("a,b;c", ",;")        // "abc"
```

#### `ReplaceNewLines(s, repl string) string`

Replaces end-of-line markers (`\r\n` and `\n`, see [`bufio.ScanLines`](https://pkg.go.dev/bufio#ScanLines)) within `s` with `repl`.
CRLF markers are matched before LF, so they are consumed whole and never corrupted into `\r` + `repl`.

```go
stringhelper.ReplaceNewLines("a\r\nb\nc", " ")  // "a b c"
```

### Joining

#### `JoinSkip(elems []string, sep string, skip func(string) bool) string`

Like [`strings.Join`](https://pkg.go.dev/strings#Join), but skips the elements for which `skip` returns true.
If `skip` is `nil`, behaves exactly like `strings.Join`.

```go
notEmpty := func(s string) bool { return s == "" }
stringhelper.JoinSkip([]string{"a", "", "b"}, "-", notEmpty)  // "a-b"
```

### Splitting

#### `StringToStrings(s string) []string`

Slices `s` into all substrings separated by end-of-line markers (see [`bufio.ScanLines`](https://pkg.go.dev/bufio#ScanLines)).
If `s` is empty, a slice with a single empty-string element is returned.

```go
stringhelper.StringToStrings("a\r\nb\nc")  // []string{"a", "b", "c"}
```

### Predicates

#### `IsEmptyOrWhite(s string) bool`

Reports whether `s` is empty or contains only white spaces, as defined by Unicode.

```go
stringhelper.IsEmptyOrWhite("")      // true
stringhelper.IsEmptyOrWhite("  \t")  // true
stringhelper.IsEmptyOrWhite("a")     // false
```

#### `IsUpper(s string) bool`

Reports whether `s` is upper case.
A string containing no cased runes (e.g. `"123"`) is considered both upper and lower case.

```go
stringhelper.IsUpper("ABC")  // true
stringhelper.IsUpper("123")  // true
stringhelper.IsUpper("Abc")  // false
```

#### `IsLower(s string) bool`

Reports whether `s` is lower case.
A string containing no cased runes (e.g. `"123"`) is considered both upper and lower case.

```go
stringhelper.IsLower("abc")  // true
stringhelper.IsLower("123")  // true
stringhelper.IsLower("Abc")  // false
```
