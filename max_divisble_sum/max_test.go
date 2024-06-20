package main

import "testing"

func Test_maxSum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{arr: []int{3, 6, 5, 1, 8}},
			want: 18,
		},
		{
			name: "test case 2",
			args: args{arr: []int{4}},
			want: 0,
		},
		{
			name: "test case 3",
			args: args{arr: []int{1, 2, 3, 4, 4}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSum(tt.args.arr); got != tt.want {
				t.Errorf("maxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
