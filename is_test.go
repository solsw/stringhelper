package stringhelper

import (
	"testing"
)

func TestIsEmptyOrWhite(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{s: ""}, want: true},
		{name: "2", args: args{s: "  "}, want: true},
		{name: "3", args: args{s: "\t"}, want: true},
		{name: "4", args: args{s: "\t \n"}, want: true},
		{name: "5", args: args{s: "qwerty"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmptyOrWhite(tt.args.s); got != tt.want {
				t.Errorf("IsEmptyOrWhite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsUpper(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{s: "QWERTY"}, want: true},
		{name: "2", args: args{s: "Qwerty"}, want: false},
		{name: "3", args: args{s: "qwertY"}, want: false},
		{name: "4", args: args{s: "asdfgh"}, want: false},
		{name: "5", args: args{s: "Б"}, want: true},
		{name: "6", args: args{s: "ф"}, want: false},
		{name: "7", args: args{s: "01234"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUpper(tt.args.s); got != tt.want {
				t.Errorf("IsUpper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "1", args: args{s: "QWERTY"}, want: false},
		{name: "2", args: args{s: "Qwerty"}, want: false},
		{name: "3", args: args{s: "qwertY"}, want: false},
		{name: "4", args: args{s: "asdfgh"}, want: true},
		{name: "5", args: args{s: "Б"}, want: false},
		{name: "6", args: args{s: "ф"}, want: true},
		{name: "7", args: args{s: "01234"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLower(tt.args.s); got != tt.want {
				t.Errorf("IsLower() = %v, want %v", got, tt.want)
			}
		})
	}
}
