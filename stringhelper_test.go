package stringhelper

import (
	"reflect"
	"testing"
)

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

func TestStringToStrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "01", args: args{s: ""}, want: []string{}},
		{name: "02", args: args{s: " "}, want: []string{" "}},
		{name: "1", args: args{s: "2\n3\r\n4"}, want: []string{"2", "3", "4"}},
		{name: "2", args: args{s: "1\r\n\n3\r\n4"}, want: []string{"1", "", "3", "4"}},
		{name: "3", args: args{s: "2\n3\r\n4\n"}, want: []string{"2", "3", "4"}},
		{name: "4", args: args{s: "\r\n2\n3\r\n4\n"}, want: []string{"", "2", "3", "4"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StringToStrings(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringToStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
