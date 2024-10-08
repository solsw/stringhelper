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
		s   string
		new string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "00",
			args: args{
				s:   "",
				new: "",
			},
			want: "",
		},
		{name: "01",
			args: args{
				s:   "",
				new: "b",
			},
			want: "",
		},
		{name: "02",
			args: args{
				s:   "\n",
				new: "a",
			},
			want: "a",
		},
		{name: "03",
			args: args{
				s:   "\r\n",
				new: "c",
			},
			want: "c",
		},
		{name: "04",
			args: args{
				s:   "\n\r\n",
				new: "d",
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
				s:   "2\r\n3",
				new: "e",
			},
			want: "2e3",
		},
		{name: "4",
			args: args{
				s:   "2\r\n3\n4",
				new: "f",
			},
			want: "2f3f4",
		},
		{name: "5",
			args: args{
				s:   "2\n3\r\n4",
				new: "g",
			},
			want: "2g3g4",
		},
		{name: "6",
			args: args{
				s:   "2\n\r\n3\r\n\n4",
				new: "h",
			},
			want: "2hh3hh4",
		},
		{name: "7",
			args: args{
				s:   "2\n3\n",
				new: "i",
			},
			want: "2i3i",
		},
		{name: "8",
			args: args{
				s:   "2\n3\r\n",
				new: "j",
			},
			want: "2j3j",
		},
		{name: "9",
			args: args{
				s:   "\r\n2\r\n3\n\n",
				new: "k",
			},
			want: "k2k3kk",
		},
		{name: "10",
			args: args{
				s:   "2\n3\r\n4\n",
				new: "l",
			},
			want: "2l3l4l",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceNewLines(tt.args.s, tt.args.new); got != tt.want {
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
