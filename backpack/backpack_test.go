package backpack

import (
	"testing"
)

func Test_bakcPack01(t *testing.T) {
	type args struct {
		m int
		w []int
		v []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				m: 4,
				w: []int{1, 2, 3},
				v: []int{6, 10, 12},
			},
			want: 18,
		},
		{
			name: "test case 2",
			args: args{
				m: 5,
				w: []int{1, 2, 3},
				v: []int{6, 10, 12},
			},
			want: 22,
		},
		{
			name: "test case 3",
			args: args{
				m: 100,
				w: []int{10, 20, 30, 40, 50},
				v: []int{20, 30, 66, 40, 60},
			},
			want: 156,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bakcPack01(tt.args.m, tt.args.w, tt.args.v); got != tt.want {
				t.Errorf("bakcPack01() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_backPackFull(t *testing.T) {
	type args struct {
		m int
		w []int
		v []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				m: 4,
				w: []int{1, 2, 3},
				v: []int{6, 10, 12},
			},
			want: 24,
		},
		{
			name: "test case 2",
			args: args{
				m: 5,
				w: []int{1, 2, 3},
				v: []int{4, 10, 12},
			},
			want: 24,
		},
		{
			name: "test case 3",
			args: args{
				m: 100,
				w: []int{10, 20, 30, 40, 50},
				v: []int{20, 30, 66, 40, 60},
			},
			want: 218,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := backPackFull(tt.args.m, tt.args.w, tt.args.v); got != tt.want {
				t.Errorf("backPackFull() = %v, want %v", got, tt.want)
			}
		})
	}
}
