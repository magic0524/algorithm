package numsum

import "testing"

func Test_min(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{a: 12},
			want: 3,
		},
		{
			name: "test case 1",
			args: args{a: 15},
			want: 4,
		},
		{
			name: "test case 1",
			args: args{a: 30},
			want: 3,
		},
		{
			name: "test case 1",
			args: args{a: 25},
			want: 1,
		},
		{
			name: "test case 1",
			args: args{a: 120},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.a); got != tt.want {
				t.Errorf("numSun() = %v, want %v", got, tt.want)
			}
		})
	}
}
