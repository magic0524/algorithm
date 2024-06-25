package stock

import "testing"

func Test_stockMax(t *testing.T) {
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
			args: args{arr: []int{7, 1, 5, 3, 6, 4}},
			want: 7,
		},
		{
			name: "test case 2",
			args: args{arr: []int{1, 2, 3, 4, 5}},
			want: 4,
		},
		{
			name: "test case 3",
			args: args{arr: []int{7, 6, 4, 3, 1}},
			want: 0,
		},
		{
			name: "test case 4",
			args: args{arr: []int{7, 6, 7, 3, 1}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stockMax(tt.args.arr); got != tt.want {
				t.Errorf("stockMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
