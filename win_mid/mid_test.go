package winmid

import (
	"reflect"
	"testing"
)

func Test_winMid(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case 1",
			args: args{arr: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 3},
			want: []int{1, -1, -1, 3, 5, 6},
		},
		{
			name: "test case 2",
			args: args{arr: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 4},
			want: []int{0, 1, 1, 4, 5},
		},
		{
			name: "test case 3",
			args: args{arr: []int{1, 3, -1, -3, 5, 3, 6, 7}, k: 5},
			want: []int{1, 3, 3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winMid(tt.args.arr, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("winMid() = %v, want %v", got, tt.want)
			}
		})
	}
}
