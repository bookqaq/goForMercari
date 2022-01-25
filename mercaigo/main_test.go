package mercarigo

import (
	"net/url"
	"reflect"
	"testing"
)

func Test_fetch(t *testing.T) {
	type args struct {
		baseURL string
		data    searchData
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{"test1",
			args{
				baseURL: "https://api.mercari.jp/search_index/search",
				data:    searchData{Keyword: "sasakure", Limit: 30, Page: 0, Sort: "created_time", Order: "desc", Status: "on_sale"}},
			""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fetch(tt.args.baseURL, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchData_paramGet(t *testing.T) {
	tests := []struct {
		name string
		data *searchData
		want url.Values
	}{
		{"test1", &searchData{Keyword: "sasakure", Limit: 30, Page: 0, Sort: "created_time", Order: "desc", Status: "on_sale"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.data.paramGet(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchData.paramGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
