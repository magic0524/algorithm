package min

import "testing"

func Test_minEditDistance(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{"horse", "ros"},
			want: 3,
		},
		{
			name: "test case 2",
			args: args{"hello", "world"},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minEditDistance(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("minEditDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
