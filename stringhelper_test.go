package stringhelper

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	type args struct {
		s  string
		i  int
		ss []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "all empty",
			want: "",
		},
		{name: "1234",
			args: args{
				s:  "1234",
				i:  1,
				ss: []string{"5", "6"},
			},
			want: "156234",
		},
		{name: "йцукен",
			args: args{
				s:  "йцукен",
				i:  1,
				ss: []string{"фывапр"},
			},
			want: "йфывапрцукен",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Insert(tt.args.s, tt.args.i, tt.args.ss...); got != tt.want {
				t.Errorf("Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertPanic(t *testing.T) {
	type args struct {
		s  string
		i  int
		ss []string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "negative index",
			args: args{s: "1234", i: -1, ss: []string{"x"}},
		},
		{name: "index past end",
			args: args{s: "1234", i: 5, ss: []string{"x"}},
		},
		{name: "rune index past end",
			args: args{s: "йцукен", i: 7, ss: []string{"x"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Errorf("Insert() did not panic for i = %d", tt.args.i)
				}
			}()
			_ = Insert(tt.args.s, tt.args.i, tt.args.ss...)
		})
	}
}

func TestSkipAny(t *testing.T) {
	type args struct {
		s     string
		chars string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: args{
				s:     "qwerty",
				chars: "yq",
			},
			want: "wert",
		},
		{name: "2",
			args: args{
				s:     "qwerty",
				chars: "ytrewq",
			},
			want: "",
		},
		{name: "3",
			args: args{
				s:     "qwerty",
				chars: "asdfgh",
			},
			want: "qwerty",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SkipAny(tt.args.s, tt.args.chars); got != tt.want {
				t.Errorf("SkipAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJoinSkip(t *testing.T) {
	type args struct {
		elems []string
		sep   string
		skip  func(string) bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "1",
			args: args{
				elems: []string{"foo", "bar", "baz"},
				sep:   ", ",
				skip:  nil,
			},
			want: "foo, bar, baz",
		},
		{name: "2",
			args: args{
				elems: []string{"foo", "", "baz"},
				sep:   ", ",
				skip:  nil,
			},
			want: "foo, , baz",
		},
		{name: "3",
			args: args{
				elems: []string{"foo", "", "baz"},
				sep:   ", ",
				skip:  func(s string) bool { return s == "" },
			},
			want: "foo, baz",
		},
		{name: "4 leading empty",
			args: args{
				elems: []string{"", "bar"},
				sep:   ", ",
				skip:  nil,
			},
			want: ", bar",
		},
		{name: "5 leading empties",
			args: args{
				elems: []string{"", "", "baz"},
				sep:   ", ",
				skip:  nil,
			},
			want: ", , baz",
		},
		{name: "6 skip leading empty",
			args: args{
				elems: []string{"", "bar", "", "baz"},
				sep:   ", ",
				skip:  func(s string) bool { return s == "" },
			},
			want: "bar, baz",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JoinSkip(tt.args.elems, tt.args.sep, tt.args.skip); got != tt.want {
				t.Errorf("JoinSkip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReplaceNewLines(t *testing.T) {
	type args struct {
		s    string
		repl string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "00",
			args: args{
				s:    "",
				repl: "",
			},
			want: "",
		},
		{name: "01",
			args: args{
				s:    "",
				repl: "b",
			},
			want: "",
		},
		{name: "02",
			args: args{
				s:    "\n",
				repl: "a",
			},
			want: "a",
		},
		{name: "03",
			args: args{
				s:    "\r\n",
				repl: "c",
			},
			want: "c",
		},
		{name: "04",
			args: args{
				s:    "\n\r\n",
				repl: "d",
			},
			want: "dd",
		},
		{name: "1",
			args: args{
				s: "23",
			},
			want: "23",
		},
		{name: "2",
			args: args{
				s: "2\r\n3",
			},
			want: "23",
		},
		{name: "3",
			args: args{
				s:    "2\r\n3",
				repl: "e",
			},
			want: "2e3",
		},
		{name: "4",
			args: args{
				s:    "2\r\n3\n4",
				repl: "f",
			},
			want: "2f3f4",
		},
		{name: "5",
			args: args{
				s:    "2\n3\r\n4",
				repl: "g",
			},
			want: "2g3g4",
		},
		{name: "6",
			args: args{
				s:    "2\n\r\n3\r\n\n4",
				repl: "h",
			},
			want: "2hh3hh4",
		},
		{name: "7",
			args: args{
				s:    "2\n3\n",
				repl: "i",
			},
			want: "2i3i",
		},
		{name: "8",
			args: args{
				s:    "2\n3\r\n",
				repl: "j",
			},
			want: "2j3j",
		},
		{name: "9",
			args: args{
				s:    "\r\n2\r\n3\n\n",
				repl: "k",
			},
			want: "k2k3kk",
		},
		{name: "10",
			args: args{
				s:    "2\n3\r\n4\n",
				repl: "l",
			},
			want: "2l3l4l",
		},
		{name: "crlf bare lf",
			args: args{
				s:    "2\n3",
				repl: "\r\n",
			},
			want: "2\r\n3",
		},
		{name: "crlf existing crlf preserved",
			args: args{
				s:    "2\r\n3",
				repl: "\r\n",
			},
			want: "2\r\n3",
		},
		{name: "crlf mixed",
			args: args{
				s:    "2\n\r\n3",
				repl: "\r\n",
			},
			want: "2\r\n\r\n3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceNewLines(tt.args.s, tt.args.repl); got != tt.want {
				t.Errorf("ReplaceNewLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringToStrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "01",
			args: args{s: ""},
			want: []string{""},
		},
		{name: "02",
			args: args{s: " "},
			want: []string{" "},
		},
		{name: "1",
			args: args{s: "2\n3\r\n4"},
			want: []string{"2", "3", "4"},
		},
		{name: "2",
			args: args{s: "1\r\n\n3\r\n4"},
			want: []string{"1", "", "3", "4"},
		},
		{name: "3",
			args: args{s: "2\n3\r\n4\n"},
			want: []string{"2", "3", "4", ""},
		},
		{name: "4",
			args: args{s: "\r\n2\n3\r\n4\n"},
			want: []string{"", "2", "3", "4", ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToStrings(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
