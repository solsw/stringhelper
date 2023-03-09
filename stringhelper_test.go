package stringhelper

import (
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
