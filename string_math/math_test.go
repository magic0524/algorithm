package stringmath

import "testing"

func Test_math(t *testing.T) {
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
			args: args{a: "rabbbit", b: "rabbit"},
			want: 3,
		},
		{
			name: "test case 2",
			args: args{a: "rabbbit", b: "rabbbit"},
			want: 1,
		},
		{
			name: "test case 3",
			args: args{a: "rabbbit", b: "rabit"},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := math(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("math() = %v, want %v", got, tt.want)
			}
		})
	}
}
