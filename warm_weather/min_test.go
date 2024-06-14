package warmweather

import (
	"reflect"
	"testing"
)

func Test_warmWeather(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case 1",
			args: args{a: []int{1, 2, 3, 4, 5}},
			want: []int{1, 1, 1, 1, 0},
		},
		{
			name: "test case 2",
			args: args{a: []int{3, 2, 3, 4, 5}},
			want: []int{3, 1, 1, 1, 0},
		},
		{
			name: "test case 3",
			args: args{a: []int{5, 2, 4, 4, 5}},
			want: []int{0, 1, 2, 1, 0},
		},
		{
			name: "test case 4",
			args: args{a: []int{5, 2, 1, 4, 5}},
			want: []int{0, 2, 1, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := warmWeather(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("warnWeather() = %v, want %v", got, tt.want)
			}
		})
	}
}
