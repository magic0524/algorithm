package max

import (
	"testing"
)

func Test_findLargestSumSubarray(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test case 1",
			args: args{arr: []int{1, 2, 3, 4, 5}, k: 2},
			want: 15,
		},
		{
			name: "Test case 2",
			args: args{arr: []int{1, -1, 3, 4, 5}, k: 2},
			want: 13,
		},
		{
			name: "Test case 3",
			args: args{arr: []int{1, -1, -3, 4, 5}, k: 2},
			want: 9,
		},
		{
			name: "Test case 3",
			args: args{arr: []int{1, -1, -3, 4, 5}, k: 3},
			want: 10,
		},
		{
			name: "Test case 4",
			args: args{arr: []int{1, -1, 0, 4, 5}, k: 2},
			want: 10,
		},
		{
			name: "Test case 5",
			args: args{arr: []int{-1, -2, -3}, k: 1},
			want: -1,
		},
		{
			name: "Test case 6",
			args: args{arr: []int{1, -1, 0, 4, 5}, k: 1},
			want: 9,
		},
		{
			name: "Test case 7",
			args: args{arr: []int{-3, -2, -1}, k: 1},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLargestSumSubarray(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("findLargestSumSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
