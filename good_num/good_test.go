package goodnum

import "testing"

func Test_findGood(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1",
			args: args{
				arr: []int{23, 2, 4, 6, 7},
				k:   6,
			},
			want: true,
		},
		{
			name: "test case 2",
			args: args{
				arr: []int{23, 2, 4, 6, 7},
				k:   10,
			},
			want: true,
		},
		{
			name: "test case 3",
			args: args{
				arr: []int{23, 2, 4, 6, 7},
				k:   11,
			},
			want: false,
		},
		{
			name: "test case 4",
			args: args{
				arr: []int{23, 2, 4, 6, 7},
				k:   12,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findGood(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("findGood() = %v, want %v", got, tt.want)
			}
		})
	}
}
