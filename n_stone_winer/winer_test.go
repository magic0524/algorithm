package winer

import "testing"

func Test_winer_3(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{n: 4},
			want: 0,
		},
		{
			name: "test case 2",
			args: args{n: 5},
			want: 1,
		},
		{
			name: "test case 3",
			args: args{n: 6},
			want: 1,
		},
		{
			name: "test case 4",
			args: args{n: 18},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := winer_3(tt.args.n); got != tt.want {
				t.Errorf("winer_3() = %v, want %v", got, tt.want)
			}
		})
	}
}
