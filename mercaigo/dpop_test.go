package mercarigo

import (
	"testing"
)

//pass
func Test_intToBase64URL(t *testing.T) {
	type args struct {
		target int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"46", args{46}, "Lg=="},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToBase64URL(tt.args.target); got != tt.want {
				t.Errorf("intToBase64URL() = %v, want %v", got, tt.want)
			}
		})
	}
}

//pass
func Test_stringToBase64URL(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"mercari", args{"mercari"}, "bWVyY2FyaQ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringToBase64URL(tt.args.target); got != tt.want {
				t.Errorf("stringToBase64URL() = %v, want %v", got, tt.want)
			}
		})
	}
}
