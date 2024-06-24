package binary

import (
	"fmt"
	"testing"
)

func Test_binarySearch(t *testing.T) {
	fmt.Printf("====================================\n")
	type args struct {
		a []int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{a: []int{3, 4, 5, 1, 2}, k: 3},
			want: 0,
		},
		{
			name: "test case 2",
			args: args{a: []int{3, 4, 5, 1, 2}, k: -1},
			want: -1,
		},
		{
			name: "test case 3",
			args: args{a: []int{3, 4}, k: -1},
			want: -1,
		},
		{
			name: "test case 4",
			args: args{a: []int{3, 4}, k: 4},
			want: 1,
		},
		{
			name: "test case 5",
			args: args{a: []int{3}, k: 3},
			want: 0,
		},
		{
			name: "test case 6",
			args: args{a: []int{3}, k: 4},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.a, tt.args.k); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_binarySearchMin(t *testing.T) {
	fmt.Printf("====================================\n")
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{a: []int{3, 4, 5, 1, 2}},
			want: 1,
		},
		{
			name: "test case 2",
			args: args{a: []int{6, 7, 8, 1, 2, 3, 4, 5}},
			want: 1,
		},
		{
			name: "test case 3",
			args: args{a: []int{6, 7, 8, 2, 3, 4, 5}},
			want: 2,
		},
		{
			name: "test case 4",
			args: args{a: []int{6, 7}},
			want: 6,
		},
		{
			name: "test case 5",
			args: args{a: []int{10, 13, 15, 3, 5, 8}},
			want: 3,
		},
		{
			name: "test case 6",
			args: args{a: []int{10, 13, 15, 3}},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearchMin(tt.args.a); got != tt.want {
				t.Errorf("binV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
