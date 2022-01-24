package mercarigo

import "testing"

func Test_searchData_paramGet(t *testing.T) {
	tests := []struct {
		name string
		data *searchData
		want string
	}{
		//{"test1", &searchData{Keyword: "sasakure"}}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.data.paramGet(); got != tt.want {
				t.Errorf("searchData.paramGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
