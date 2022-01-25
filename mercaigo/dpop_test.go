package mercarigo

import (
	"reflect"
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

func Test_dPoPGenerator(t *testing.T) {
	type args struct {
		uuid_  string
		method string
		url_   string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"A must-fail test", args{"bookqaq", "GET", "https://api.mercari.jp/search_index/search"}, make([]byte, 0)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dPoPGenerator(tt.args.uuid_, tt.args.method, tt.args.url_); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dPoPGenerator() = %s, want %v", got, tt.want)
			}
		})
	}
}
