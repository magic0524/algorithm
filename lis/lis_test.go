package main

import "testing"

func Test_lis(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testcase 1",
			args: args{a: []int{10, 22, 9, 33, 21, 50, 41, 60, 80}},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lis(tt.args.a); got != tt.want {
				t.Errorf("lis() = %v, want %v", got, tt.want)
			}
		})
	}
}
